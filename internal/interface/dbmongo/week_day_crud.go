package dbmongo

import (
	"context"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDBWeekDayCrud interface {
	Find(context.Context) ([]domain.WeekDay, error)
}

type ImpDBWeekDayCrud struct {
	*mongo.Client
}

func NewDBWeekDayCrud(cli *mongo.Client) IDBWeekDayCrud {
	return &ImpDBWeekDayCrud{
		cli,
	}
}

func (wd *ImpDBWeekDayCrud) Find(ctx context.Context) ([]domain.WeekDay, error) {
	var ret []domain.WeekDay

	findOpt := options.Find()
	collRef := wd.Database("fitness").Collection("week_days")
	cursor, err := collRef.Find(ctx, bson.D{}, findOpt)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var wd domain.WeekDay

		if err := cursor.Decode(&wd); err != nil {
			panic(err)
		}

		ret = append(ret, wd)
	}

	return ret, nil
}
