package dbmongo

import (
	"context"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDbCategoryCrud interface {
	GetAllCategories(context.Context) ([]domain.Category, error)
	InsertCategory(context.Context, domain.Category) error
}

type ImpDbCategoryCrud struct {
	*mongo.Client
}

func NewDbCategoryCrud(cli *mongo.Client) IDbCategoryCrud {
	return &ImpDbCategoryCrud{
		cli,
	}
}

func (cc *ImpDbCategoryCrud) GetAllCategories(c context.Context) ([]domain.Category, error) {
	var cs []domain.Category

	findOpt := options.Find()
	docRef := cc.Client.Database("fitness").Collection("categories")
	cursor, err := docRef.Find(c, bson.D{{}}, findOpt)

	if err != nil {
		return nil, err
	}

	for cursor.Next(c) {
		var r domain.Category

		if err := cursor.Decode(&r); err != nil {
			panic(err)
		}

		cs = append(cs, r)
	}

	return cs, nil
}
func (cc *ImpDbCategoryCrud) InsertCategory(c context.Context, ses domain.Category) error {
	docRef := cc.Client.Database("fitness").Collection("categories")

	_, err := docRef.InsertOne(c, ses)

	if err != nil {
		return err
	}

	return nil
}
