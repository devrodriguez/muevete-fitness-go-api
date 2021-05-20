package weeklies

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/gin-gonic/gin"
)

type IWeeklyCrud interface {
	CreateWeekly(c *gin.Context, wk domain.Weekly) error
	GetAllWeeklies(c *gin.Context) ([]domain.Weekly, error)
}

type ImpWeeklyCrud struct {
	dbImp dbmongo.IDbWeeklyCrud
}

func NewCustomerCrud(dbImp dbmongo.IDbWeeklyCrud) IWeeklyCrud {
	return &ImpWeeklyCrud{
		dbImp,
	}
}

func (wc *ImpWeeklyCrud) CreateWeekly(c *gin.Context, wk domain.Weekly) error {
	err := wc.dbImp.InsertWeekly(c, wk)

	if err != nil {
		return err
	}

	return nil
}

func (wc *ImpWeeklyCrud) GetAllWeeklies(c *gin.Context) ([]domain.Weekly, error) {
	wks, err := wc.dbImp.FindWeekly(c)

	if err != nil {
		return nil, err
	}

	return wks, nil
}
