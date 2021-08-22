package dbmongo

import (
	"context"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDbWeeklyCrud interface {
	SaveWeekly(context.Context, domain.WeeklyMod) error
	FindWeekly(context.Context) ([]domain.Weekly, error)
}

type ImpDbWeeklyCrud struct {
	*mongo.Client
}

func NewDbWeeklyCrud(cli *mongo.Client) IDbWeeklyCrud {
	return &ImpDbWeeklyCrud{
		cli,
	}
}

func (re *ImpDbWeeklyCrud) FindWeekly(ctx context.Context) ([]domain.Weekly, error) {
	var wks []domain.Weekly

	docRef := re.Client.Database("fitness").Collection("weeklies")

	lookSes := bson.D{
		{"$lookup", bson.D{
			{"from", "sessions"},
			{"localField", "session"},
			{"foreignField", "_id"},
			{"as", "session"},
		}}}

	unwindSes := bson.D{
		{"$unwind", bson.D{
			{"path", "$session"},
			{"preserveNullAndEmptyArrays", false},
		}},
	}

	lookRouSch := bson.D{
		{"$lookup", bson.D{
			{"from", "routine_schedule"},
			{"localField", "routine_schedule"},
			{"foreignField", "_id"},
			{"as", "routine_schedule"},
		}},
	}

	unwindRouSch := bson.D{
		{"$unwind", bson.D{
			{"path", "$routine_schedule"},
			{"preserveNullAndEmptyArrays", false},
		}},
	}

	lookRouSchRou := bson.D{
		{"$lookup", bson.D{
			{"from", "routines"},
			{"localField", "routine_schedule.routine"},
			{"foreignField", "_id"},
			{"as", "routine_schedule.routine"},
		}},
	}

	unwindRouSchRou := bson.D{
		{"$unwind", bson.D{
			{"path", "$routine_schedule.routine"},
			{"preserveNullAndEmptyArrays", false},
		}},
	}

	lookRouSchWkd := bson.D{
		{"$lookup", bson.D{
			{"from", "week_days"},
			{"localField", "routine_schedule.week_day"},
			{"foreignField", "_id"},
			{"as", "routine_schedule.week_day"},
		}},
	}

	unwindRouSchWkd := bson.D{
		{"$unwind", bson.D{
			{"path", "$routine_schedule.week_day"},
			{"preserveNullAndEmptyArrays", false},
		}},
	}

	cursor, err := docRef.Aggregate(ctx, mongo.Pipeline{
		lookSes,
		unwindSes,
		lookRouSch,
		unwindRouSch,
		lookRouSchRou,
		unwindRouSchRou,
		lookRouSchWkd,
		unwindRouSchWkd,
	})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var wk domain.Weekly

		if err := cursor.Decode(&wk); err != nil {
			panic(err)
		}

		wks = append(wks, wk)
	}

	return wks, nil
}

func (re *ImpDbWeeklyCrud) SaveWeekly(ctx context.Context, wk domain.WeeklyMod) error {
	collRef := re.Client.Database("fitness").Collection("weeklies")

	_, err := collRef.InsertOne(ctx, wk)

	if err != nil {
		return err
	}

	return nil
}
