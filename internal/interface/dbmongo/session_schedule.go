package dbmongo

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDbSessionSchedule interface {
	GetSessionSchedule(c *gin.Context) ([]domain.ScheduleSession, error)
}

type ImpDbSessionSchedule struct {
	*mongo.Client
}

func NewDbSessionSchedule(cli *mongo.Client) IDbSessionSchedule {
	return &ImpDbSessionSchedule{
		cli,
	}
}

func (ss *ImpDbSessionSchedule) GetSessionSchedule(c *gin.Context) ([]domain.ScheduleSession, error) {
	var sess []domain.ScheduleSession

	//findOpt := options.Find()
	docRef := ss.Client.Database("fitness").Collection("routine_schedule")
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
	cursor, err := docRef.Aggregate(c, mongo.Pipeline{lookCust, unwindCust, lookWeek, unwindWeek, lookSes, unwindSes})
	// cursor, err := docRef.Find(c, bson.D{{}}, findOpt)

	if err != nil {
		return nil, err
	}

	for cursor.Next(c) {
		var ses domain.ScheduleSession

		if err := cursor.Decode(&ses); err != nil {
			panic(err)
		}

		sess = append(sess, ses)
	}

	return sess, nil
}

