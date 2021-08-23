package customers

import (
	"context"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
)

type ICustomerCrud interface {
	GetAllCustomers(c context.Context) ([]domain.Customer, error)
	CreateCustomer(c context.Context, ds domain.Customer) error
}

type ImpCrudCustomer struct {
	dbImp dbmongo.IDbCustomerCrud
}

func NewCustomerCrud(dbImp dbmongo.IDbCustomerCrud) ICustomerCrud {
	return &ImpCrudCustomer{
		dbImp,
	}
}

func (cs *ImpCrudCustomer) GetAllCustomers(c context.Context) ([]domain.Customer, error) {
	ses, err := cs.dbImp.GetAllCustomers(c)

	if err != nil {
		return nil, err
	}

	return ses, nil
}

func (cs *ImpCrudCustomer) CreateCustomer(c context.Context, ses domain.Customer) error {
	err := cs.dbImp.InsertCustomer(c, ses)

	if err != nil {
		return err
	}

	return nil
}
