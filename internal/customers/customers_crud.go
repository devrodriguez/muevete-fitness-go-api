package customers

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/gin-gonic/gin"
)

type ICustomerCrud interface {
	GetAllCustomers(c *gin.Context) ([]domain.Customer, error)
	CreateCustomer(c *gin.Context, ds domain.Customer) error
}

type ImpCrudCustomer struct {
	dbImp dbmongo.IDbCustomerCrud
}

func NewCustomerCrud(dbImp dbmongo.IDbCustomerCrud) ICustomerCrud {
	return &ImpCrudCustomer{
		dbImp,
	}
}

func (cs *ImpCrudCustomer) GetAllCustomers(c *gin.Context) ([]domain.Customer, error) {
	ses, err := cs.dbImp.GetAllCustomers(c)

	if err != nil {
		return nil, err
	}

	return ses, nil
}

func (cs *ImpCrudCustomer) CreateCustomer(c *gin.Context, ses domain.Customer) error {
	err := cs.dbImp.CreateCustomer(c, ses)

	if err != nil {
		return err
	}

	return nil
}
