package routines

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/gin-gonic/gin"
)

type ICrudRoutine interface {
	GetAll(c *gin.Context) ([]domain.Routine, error)
	Create(c *gin.Context, ds domain.Routine) error
}

type ImpCrudRoutine struct {
	dbImp dbmongo.IDbRoutineCrud
}

func NewCrudRoutine(dbImp dbmongo.IDbRoutineCrud) ICrudRoutine {
	return &ImpCrudRoutine{
		dbImp,
	}
}

func (cs *ImpCrudRoutine) GetAll(c *gin.Context) ([]domain.Routine, error) {
	rs, err := cs.dbImp.GetAll(c)

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cs *ImpCrudRoutine) Create(c *gin.Context, r domain.Routine) error {
	err := cs.dbImp.Create(c, r)

	if err != nil {
		return err
	}

	return nil
}

