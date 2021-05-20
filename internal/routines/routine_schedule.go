package routines

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/gin-gonic/gin"
)

type IRoutineSchedule interface {
	CreateSchedule(c *gin.Context, ds domain.RoutineSchedule) error
}

type ImpRoutineSchedule struct {
	dbImp dbmongo.IDbRoutineSchedule
}

func NewRoutineSchedule(dbImp dbmongo.IDbRoutineSchedule) IRoutineSchedule {
	return &ImpRoutineSchedule{
		dbImp,
	}
}

func (cs *ImpRoutineSchedule) CreateSchedule(c *gin.Context, sch domain.RoutineSchedule) error {
	err := cs.dbImp.InsertSchedule(c, sch)

	if err != nil {
		return err
	}

	return nil
}
