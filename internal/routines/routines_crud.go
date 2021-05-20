package routines

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/gin-gonic/gin"
)

type ICrudRoutine interface {
	GetAllRoutines(c *gin.Context) ([]domain.Routine, error)
	CreateRoutine(c *gin.Context, ds domain.Routine) error
}

type ImpCrudRoutine struct {
	dbImp dbmongo.IDbRoutineCrud
}

func NewCrudRoutine(dbImp dbmongo.IDbRoutineCrud) ICrudRoutine {
	return &ImpCrudRoutine{
		dbImp,
	}
}

func (cs *ImpCrudRoutine) GetAllRoutines(c *gin.Context) ([]domain.Routine, error) {
	rs, err := cs.dbImp.GetAllRoutines(c)

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cs *ImpCrudRoutine) CreateRoutine(c *gin.Context, r domain.Routine) error {
	err := cs.dbImp.InsertRoutine(c, r)

	if err != nil {
		return err
	}

	return nil
}
