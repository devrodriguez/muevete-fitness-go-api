package weeklies

import (
	"context"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
)

type IWeekDayCrud interface {
	GetAllDays(context.Context) ([]domain.WeekDay, error)
	SaveWeekDay(context.Context, domain.WeekDay) (*domain.WeekDay, error)
}

type ImpWeekDayCrud struct {
	dbImp dbmongo.IDBWeekDayCrud
}

func NewWeekDayCrud(dbImp dbmongo.IDBWeekDayCrud) IWeekDayCrud {
	return &ImpWeekDayCrud{
		dbImp,
	}
}

func (wd *ImpWeekDayCrud) GetAllDays(ctx context.Context) ([]domain.WeekDay, error) {
	wks, err := wd.dbImp.FindWeekDay(ctx)

	if err != nil {
		return nil, err
	}

	return wks, nil
}

func (wd *ImpWeekDayCrud) SaveWeekDay(ctx context.Context, w domain.WeekDay) (*domain.WeekDay, error) {
	wkd, err := wd.dbImp.SaveWeekDay(ctx, w)
	if err != nil {
		return nil, err
	}

	return wkd, nil
}
