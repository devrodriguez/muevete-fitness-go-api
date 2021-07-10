package dbmongo

import (
	"context"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDbCustomerCrud interface {
	GetAllCustomers(context.Context) ([]domain.Customer, error)
	InsertCustomer(context.Context, domain.Customer) error
}

type ImpDbCustomerCrud struct {
	*mongo.Client
}

func NewDbCustomerCrud(cli *mongo.Client) IDbCustomerCrud {
	return &ImpDbCustomerCrud{
		cli,
	}
}

func (cc *ImpDbCustomerCrud) GetAllCustomers(c context.Context) ([]domain.Customer, error) {
	var cs []domain.Customer

	findOpt := options.Find()
	docRef := cc.Client.Database("fitness").Collection("customers")
	cursor, err := docRef.Find(c, bson.D{{}}, findOpt)

	if err != nil {
		return nil, err
	}

	for cursor.Next(c) {
		var r domain.Customer

		if err := cursor.Decode(&r); err != nil {
			panic(err)
		}

		cs = append(cs, r)
	}

	return cs, nil
}
func (cc *ImpDbCustomerCrud) InsertCustomer(c context.Context, ses domain.Customer) error {
	docRef := cc.Client.Database("fitness").Collection("customers")

	_, err := docRef.InsertOne(c, ses)

	if err != nil {
		return err
	}

	return nil
}
