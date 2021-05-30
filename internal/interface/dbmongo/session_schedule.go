package dbmongo

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDbSessionSchedule interface {
	GetSessionSchedule(c *gin.Context) ([]domain.SessionSchedule, error)
	SaveSessionSchedule(c * gin.Context, session domain.SessionScheduleMod) error
}

type ImpDbSessionSchedule struct {
	*mongo.Client
}

func NewDbSessionSchedule(cli *mongo.Client) IDbSessionSchedule {
	return &ImpDbSessionSchedule{
		cli,
	}
}

func (ss *ImpDbSessionSchedule) GetSessionSchedule(c *gin.Context) ([]domain.SessionSchedule, error) {
	var sess []domain.SessionSchedule

	docRef := ss.Client.Database("fitness").Collection("session_schedule")
	lookCust := bson.D{
		{"$lookup", bson.D{
			{"from", "customers"},
			{"localField", "customer"},
			{"foreignField", "_id"},
			{"as", "customer"},
		}}}
	unwindCust := bson.D{
		{"$unwind", bson.D{
			{"path", "$customer"},
			{"preserveNullAndEmptyArrays", false},
		}}}

	lookWeek := bson.D{
		{"$lookup", bson.D{
			{"from", "weeklies"},
			{"localField", "weekly"},
			{"foreignField", "_id"},
			{"as", "weekly"},
		}}}
	unwindWeek := bson.D{
		{"$unwind", bson.D{
			{"path", "$weekly"},
			{"preserveNullAndEmptyArrays", false},
		}}}

	lookSes := bson.D{
		{"$lookup", bson.D{
			{"from", "sessions"},
			{"localField", "weekly.session"},
			{"foreignField", "_id"},
			{"as", "weekly.session"},
		}}}
	unwindSes := bson.D{
		{"$unwind", bson.D{
			{"path", "$weekly.session"},
			{"preserveNullAndEmptyArrays", false},
		}}}

	lookRouSch := bson.D{
		{"$lookup", bson.D{
			{"from", "routine_schedule"},
			{"localField", "weekly.routine_schedule"},
			{"foreignField", "_id"},
			{"as", "weekly.routine_schedule"},
		}}}
	unwindRouSch := bson.D{
		{"$unwind", bson.D{
			{"path", "$weekly.routine_schedule"},
			{"preserveNullAndEmptyArrays", false},
		}}}

	lookRouSchRou := bson.D{
		{"$lookup", bson.D{
			{"from", "routines"},
			{"localField", "weekly.routine_schedule.routine"},
			{"foreignField", "_id"},
			{"as", "weekly.routine_schedule.routine"},
		}}}
	unwindRouSchRou := bson.D{
		{"$unwind", bson.D{
			{"path", "$weekly.routine_schedule.routine"},
			{"preserveNullAndEmptyArrays", false},
		}}}

	lookRouSchWkd := bson.D{
		{"$lookup", bson.D{
			{"from", "week_days"},
			{"localField", "weekly.routine_schedule.week_days"},
			{"foreignField", "_id"},
			{"as", "weekly.routine_schedule.week_days"},
		}}}
	unwindRouSchWkd := bson.D{
		{"$unwind", bson.D{
			{"path", "$weekly.routine_schedule.week_days"},
			{"preserveNullAndEmptyArrays", false},
		}}}

	cursor, err := docRef.Aggregate(c, mongo.Pipeline{
		lookCust,
		unwindCust,
		lookWeek,
		unwindWeek,
		lookSes,
		unwindSes,
		lookRouSch,
		unwindRouSch,
		lookRouSchRou,
		unwindRouSchRou,
		lookRouSchWkd,
		unwindRouSchWkd,
	})

	if err != nil {
		return nil, err
	}

	for cursor.Next(c) {
		var ses domain.SessionSchedule

		if err := cursor.Decode(&ses); err != nil {
			panic(err)
		}

		sess = append(sess, ses)
	}

	return sess, nil
}

func (ss *ImpDbSessionSchedule) SaveSessionSchedule(c * gin.Context, sess domain.SessionScheduleMod) error {
	docRef := ss.Client.Database("fitness").Collection("session_schedule")

	_, err := docRef.InsertOne(c, sess)

	if err != nil {
		return err
	}

	return nil
}

