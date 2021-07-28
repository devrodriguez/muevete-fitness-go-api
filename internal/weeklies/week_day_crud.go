package weeklies

import (
	"context"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
)

type IWeekDayCrud interface {
	GetAllDays(context.Context) ([]domain.WeekDay, error)
}

type ImpWeekDayCrud struct {
	dbmongo.IDBWeekDayCrud
}

func NewWeekDayCrud(dbImp dbmongo.IDBWeekDayCrud) IWeekDayCrud {
	return &ImpWeekDayCrud{
		dbImp,
	}
}

func (wd *ImpWeekDayCrud) GetAllDays(ctx context.Context) ([]domain.WeekDay, error) {
	wks, err := wd.Find(ctx)

	if err != nil {
		return nil, err
	}

	return wks, nil
}
