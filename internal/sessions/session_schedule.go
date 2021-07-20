package sessions

import (
	"context"
	"errors"
	"fmt"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
)

const SCH_CAPACITY = 20

var errCapacity = errors.New("exced scheduled capacity")
var errExist = errors.New("item exist")

type ISessionSchedule interface {
	GetSessionsSchedule(context.Context) ([]domain.SessionSchedule, error)
	CreateSessionsSchedule(context.Context, domain.SessionScheduleMod) error
}

type ImpSessionSchedule struct {
	dbImp dbmongo.IDbSessionSchedule
}

func NewSessionSchedule(dbImp dbmongo.IDbSessionSchedule) ISessionSchedule {
	return &ImpSessionSchedule{
		dbImp,
	}
}

func (cs *ImpSessionSchedule) GetSessionsSchedule(c context.Context) ([]domain.SessionSchedule, error) {
	schs, err := cs.dbImp.GetAllSessionSchedule(c)

	if err != nil {
		return nil, err
	}

	return schs, err
}

func (cs *ImpSessionSchedule) CreateSessionsSchedule(c context.Context, ss domain.SessionScheduleMod) error {
	// Validate capacity for weekly
	nDocs, err := cs.dbImp.GetByWeekly(c, ss.WeeklyID.Hex())
	if err != nil {
		return err
	}

	if nDocs > SCH_CAPACITY {
		fmt.Printf("schedule capacity excede %d", SCH_CAPACITY)
		return errCapacity
	}

	if err := cs.dbImp.SaveSessionSchedule(c, ss); err != nil {
		return err
	}

	return nil
}
