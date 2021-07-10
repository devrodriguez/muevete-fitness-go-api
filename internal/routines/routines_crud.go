package routines

import (
	"context"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
)

type ICrudRoutine interface {
	GetAllRoutines(c context.Context) ([]domain.Routine, error)
	CreateRoutine(c context.Context, ds domain.Routine) error
}

type ImpCrudRoutine struct {
	dbImp dbmongo.IDbRoutineCrud
}

func NewCrudRoutine(dbImp dbmongo.IDbRoutineCrud) ICrudRoutine {
	return &ImpCrudRoutine{
		dbImp,
	}
}

func (cs *ImpCrudRoutine) GetAllRoutines(c context.Context) ([]domain.Routine, error) {
	rs, err := cs.dbImp.GetAllRoutines(c)

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cs *ImpCrudRoutine) CreateRoutine(c context.Context, r domain.Routine) error {
	err := cs.dbImp.InsertRoutine(c, r)

	if err != nil {
		return err
	}

	return nil
}
