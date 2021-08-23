package dbmongo

import (
	"context"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDBWeekDayCrud interface {
	FindWeekDay(context.Context) ([]domain.WeekDay, error)
	SaveWeekDay(context.Context, domain.WeekDay) (*domain.WeekDay, error)
}

type ImpDBWeekDayCrud struct {
	*mongo.Client
}

func NewDBWeekDayCrud(cli *mongo.Client) IDBWeekDayCrud {
	return &ImpDBWeekDayCrud{
		cli,
	}
}

func (wd *ImpDBWeekDayCrud) FindWeekDay(ctx context.Context) ([]domain.WeekDay, error) {
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

func (wd *ImpDBWeekDayCrud) SaveWeekDay(ctx context.Context, wkd domain.WeekDay) (*domain.WeekDay, error) {
	collRef := wd.Client.Database("fitness").Collection("week_days")

	res, err := collRef.InsertOne(ctx, wkd)
	if err != nil {
		return nil, err
	}

	wkd.ID = res.InsertedID.(primitive.ObjectID)
	if err != nil {
		return nil, err
	}

	return &wkd, nil
}
