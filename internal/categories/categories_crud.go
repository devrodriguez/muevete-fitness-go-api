package categories

import (
	"context"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
)

type ICategoryCrud interface {
	GetAllCategories(context.Context) ([]domain.Category, error)
	CreateCategory(context.Context, domain.Category) error
}

type ImpCategoryCrud struct {
	dbImp dbmongo.IDbCategoryCrud
}

func NewCategoryCrud(dbImp dbmongo.IDbCategoryCrud) ICategoryCrud {
	return &ImpCategoryCrud{
		dbImp,
	}
}

func (cc *ImpCategoryCrud) GetAllCategories(c context.Context) ([]domain.Category, error) {
	ses, err := cc.dbImp.GetAllCategories(c)

	if err != nil {
		return nil, err
	}

	return ses, nil
}

func (cc *ImpCategoryCrud) CreateCategory(c context.Context, cat domain.Category) error {
	err := cc.dbImp.InsertCategory(c, cat)

	if err != nil {
		return err
	}

	return nil
}
