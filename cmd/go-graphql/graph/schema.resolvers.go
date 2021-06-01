package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/devrodriguez/muevete-fitness-go-api/cmd/go-graphql/graph/generated"
	"github.com/devrodriguez/muevete-fitness-go-api/cmd/go-graphql/graph/model"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/customers"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/interface/dbmongo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.NewCustomer) (*model.Customer, error) {
	var customer domain.Customer
	var inpCustomer model.Customer

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	customer.Name = input.Name
	customer.LastName = input.LastName
	customer.Email = input.Email

	inpCustomer.Name = input.Name
	inpCustomer.LastName = input.LastName
	inpCustomer.Email = input.Email

	cusRepo := dbmongo.NewDbCustomerCrud(client)
	cusUc := customers.NewCustomerCrud(cusRepo)

	if err := cusUc.CreateCustomer(&gin.Context{}, customer); err != nil {
		return nil, err
	}

	return &inpCustomer, nil
}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	var custs []*model.Customer

	mctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	client, err := mongo.Connect(mctx, options.Client().ApplyURI("mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/fitness?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	cusRepo := dbmongo.NewDbCustomerCrud(client)
	cusUc := customers.NewCustomerCrud(cusRepo)

	cust, _ := cusUc.GetAllCustomers(&gin.Context{})

	for _, c := range cust {
		cust := model.Customer{
			Name:     c.Name,
			LastName: c.LastName,
			Email:    c.Email,
		}
		custs = append(custs, &cust)
	}

	return custs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
