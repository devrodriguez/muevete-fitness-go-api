package routines

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/gin-gonic/gin"
)

type IRoutineSchedule interface {
	GetSchedule(*gin.Context) ([]domain.RoutineSchedule, error)
	CreateSchedule(*gin.Context, domain.RoutineScheduleMod) error
}

type ImpRoutineSchedule struct {
	dbImp dbmongo.IDbRoutineSchedule
}

func NewRoutineSchedule(dbImp dbmongo.IDbRoutineSchedule) IRoutineSchedule {
	return &ImpRoutineSchedule{
		dbImp,
	}
}

func (rs *ImpRoutineSchedule) GetSchedule(c *gin.Context) ([]domain.RoutineSchedule, error) {
	rss, err := rs.dbImp.FindRoutineSchedule(c)

	if err != nil {
		return nil, err
	}

	return rss, nil
}

func (rs *ImpRoutineSchedule) CreateSchedule(c *gin.Context, sch domain.RoutineScheduleMod) error {
	err := rs.dbImp.SaveRoutineSchedule(c, sch)

	if err != nil {
		return err
	}

	return nil
}
