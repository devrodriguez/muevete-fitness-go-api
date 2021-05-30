package dbmongo

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDbRoutineSchedule interface {
	FindRoutineSchedule(*gin.Context) ([]domain.RoutineSchedule, error)
	SaveRoutineSchedule(*gin.Context, domain.RoutineScheduleMod) error
}

type ImpDbRoutineSchedule struct {
	*mongo.Client
}

func NewDbRoutineSchedule(cli *mongo.Client) IDbRoutineSchedule {
	return &ImpDbRoutineSchedule{
		cli,
	}
}

func (re *ImpDbRoutineSchedule) FindRoutineSchedule(c *gin.Context) ([]domain.RoutineSchedule, error) {
	var rsch []domain.RoutineSchedule

	docRef := re.Client.Database("fitness").Collection("routine_schedule")

	lookRou := bson.D{
		{"$lookup", bson.D{
			{"from", "routines"},
			{"localField", "routine"},
			{"foreignField", "_id"},
			{"as", "routine"},
		}}}

	unwindRou := bson.D{
		{"$unwind", bson.D{
			{"path", "$routine"},
			{"preserveNullAndEmptyArrays", false},
		}}}

	lookDay := bson.D{
		{"$lookup", bson.D{
			{"from", "week_days"},
			{"localField", "week_days"},
			{"foreignField", "_id"},
			{"as", "week_days"},
		}}}

	unwindDay := bson.D{
		{"$unwind", bson.D{
			{"path", "$week_days"},
			{"preserveNullAndEmptyArrays", false},
		}}}

	cursor, err := docRef.Aggregate(c, mongo.Pipeline{
		lookRou,
		unwindRou,
		lookDay,
		unwindDay,
	})

	if err != nil {
		return nil, err
	}

	for cursor.Next(c) {
		var rs domain.RoutineSchedule

		if err := cursor.Decode(&rs); err != nil {
			panic(err)
		}

		rsch = append(rsch, rs)
	}

	return rsch, nil
}

func (re *ImpDbRoutineSchedule) SaveRoutineSchedule(c *gin.Context, sch domain.RoutineScheduleMod) error {
	docRef := re.Client.Database("fitness").Collection("routine_schedule")

	_, err := docRef.InsertOne(c, sch)

	if err != nil {
		return err
	}

	return nil
}
