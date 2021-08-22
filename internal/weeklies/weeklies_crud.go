package weeklies

import (
	"context"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
)

type IWeeklyCrud interface {
	CreateWeekly(context.Context, domain.WeeklyMod) error
	GetAllWeeklies(context.Context) ([]domain.Weekly, error)
}

type ImpWeeklyCrud struct {
	dbImp dbmongo.IDbWeeklyCrud
}

func NewWeeklyCrud(dbImp dbmongo.IDbWeeklyCrud) IWeeklyCrud {
	return &ImpWeeklyCrud{
		dbImp,
	}
}

func (wc *ImpWeeklyCrud) GetAllWeeklies(ctx context.Context) ([]domain.Weekly, error) {
	wks, err := wc.dbImp.FindWeekly(ctx)

	if err != nil {
		return nil, err
	}

	return wks, nil
}

func (wc *ImpWeeklyCrud) CreateWeekly(ctx context.Context, wk domain.WeeklyMod) error {
	err := wc.dbImp.SaveWeekly(ctx, wk)

	if err != nil {
		return err
	}

	return nil
}
