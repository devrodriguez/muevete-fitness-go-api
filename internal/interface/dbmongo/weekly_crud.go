package dbmongo

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDbWeeklyCrud interface {
	InsertWeekly(*gin.Context, domain.Weekly) error
	FindWeekly(*gin.Context) ([]domain.Weekly, error)
}

type ImpDbWeeklyCrud struct {
	*mongo.Client
}

func NewDbWeeklyCrud(cli *mongo.Client) IDbWeeklyCrud {
	return &ImpDbWeeklyCrud{
		cli,
	}
}

func (re *ImpDbWeeklyCrud) InsertWeekly(c *gin.Context, wk domain.Weekly) error {
	docRef := re.Client.Database("fitness").Collection("weeklies")

	_, err := docRef.InsertOne(c, wk)

	if err != nil {
		return err
	}

	return nil
}

func (re *ImpDbWeeklyCrud) FindWeekly(c *gin.Context) ([]domain.Weekly, error) {
	var wks []domain.Weekly

	findOpt := options.Find()
	docRef := re.Client.Database("fitness").Collection("weeklies")
	cursor, err := docRef.Find(c, bson.D{{}}, findOpt)

	if err != nil {
		return nil, err
	}

	for cursor.Next(c) {
		var wk domain.Weekly

		if err := cursor.Decode(&wk); err != nil {
			panic(err)
		}

		wks = append(wks, wk)
	}

	return wks, nil
}
