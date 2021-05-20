package categories

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/gin-gonic/gin"
)

type ICategoryCrud interface {
	GetAllCategories(*gin.Context) ([]domain.Category, error)
	CreateCategory(*gin.Context, domain.Category) error
}

type ImpCategoryCrud struct {
	dbImp dbmongo.IDbCategoryCrud
}

func NewCategoryCrud(dbImp dbmongo.IDbCategoryCrud) ICategoryCrud {
	return &ImpCategoryCrud{
		dbImp,
	}
}

func (cc *ImpCategoryCrud) GetAllCategories(c *gin.Context) ([]domain.Category, error) {
	ses, err := cc.dbImp.GetAllCategories(c)

	if err != nil {
		return nil, err
	}

	return ses, nil
}

func (cc *ImpCategoryCrud) CreateCategory(c *gin.Context, cat domain.Category) error {
	err := cc.dbImp.InsertCategory(c, cat)

	if err != nil {
		return err
	}

	return nil
}
