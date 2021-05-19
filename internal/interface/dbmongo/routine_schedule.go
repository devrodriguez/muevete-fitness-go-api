package dbmongo

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDbRoutineSchedule interface {
	CreateSchedule(*gin.Context, domain.RoutineSchedule) error
}

type ImpDbRoutineSchedule struct {
	*mongo.Client
}

func NewDbRoutineSchedule(cli *mongo.Client) IDbRoutineSchedule {
	return &ImpDbRoutineSchedule{
		cli,
	}
}

func (re *ImpDbRoutineSchedule) CreateSchedule(c *gin.Context, sch domain.RoutineSchedule) error {
	docRef := re.Client.Database("fitness").Collection("routine_schedule")

	_, err := docRef.InsertOne(c, sch)

	if err != nil {
		return err
	}

	return nil
}
