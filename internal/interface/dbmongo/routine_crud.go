package dbmongo

import (
	"context"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDbRoutineCrud interface {
	GetAllRoutines(context.Context) ([]domain.Routine, error)
	QRoutinesByDay(c context.Context, day string) ([]domain.RoutineCategory, error)
	InsertRoutine(context.Context, domain.Routine) error
}

type ImpDbRoutineCrud struct {
	*mongo.Client
}

func NewDbRoutineCrud(cli *mongo.Client) IDbRoutineCrud {
	return &ImpDbRoutineCrud{
		cli,
	}
}

func (rc *ImpDbRoutineCrud) GetAllRoutines(c context.Context) ([]domain.Routine, error) {
	var rs []domain.Routine

	findOpt := options.Find()
	docRef := rc.Client.Database("fitness").Collection("routines")
	cursor, err := docRef.Find(c, bson.D{{}}, findOpt)

	if err != nil {
		return nil, err
	}

	for cursor.Next(c) {
		var r domain.Routine

		if err := cursor.Decode(&r); err != nil {
			panic(err)
		}

		rs = append(rs, r)
	}

	return rs, nil
}

func (rc *ImpDbRoutineCrud) InsertRoutine(c context.Context, r domain.Routine) error {
	docRef := rc.Client.Database("fitness").Collection("routines")

	_, err := docRef.InsertOne(c, r)

	if err != nil {
		return err
	}

	return nil
}

func (rc *ImpDbRoutineCrud) QRoutinesByDay(c context.Context, day string) ([]domain.RoutineCategory, error) {
	var rcs []domain.RoutineCategory
	docRef := rc.Client.Database("fitness").Collection("routine_category")

	lookRou := bson.D{
		{"$lookup", bson.D{
			{"from", "routines"},
			{"localField", "routine_id"},
			{"foreignField", "_id"},
			{"as", "routine"},
		}}}

	unwindRou := bson.D{
		{"$unwind", bson.D{
			{"path", "$routine"},
			{"preserveNullAndEmptyArrays", false},
		}},
	}

	lookCat := bson.D{
		{"$lookup", bson.D{
			{"from", "categories"},
			{"localField", "category_id"},
			{"foreignField", "_id"},
			{"as", "category"},
		}}}

	unwindCat := bson.D{
		{"$unwind", bson.D{
			{"path", "$category"},
			{"preserveNullAndEmptyArrays", false},
		}},
	}

	groupCat := bson.D{
		{
			"$group", bson.M{
				"_id":      bson.M{"category": "$category"},
				"routines": bson.M{"$push": "$routine"},
			},
		},
	}

	groupCatLs := bson.D{
		{
			"$group", bson.M{
				"_id": nil,
				"categories": bson.M{
					"$push": bson.M{
						"category": "$_id.category",
						"routines": "$routines",
					},
				},
			},
		},
	}

	unwindCatGroup := bson.D{
		{
			"$unwind", bson.D{
				{"path", "$categories"},
				{"preserveNullAndEmptyArrays", true},
			},
		},
	}

	replaceRoot := bson.D{
		{
			"$replaceRoot", bson.D{
				{"newRoot", "$categories"},
			},
		},
	}

	cursor, err := docRef.Aggregate(c, mongo.Pipeline{
		lookRou,
		unwindRou,
		lookCat,
		unwindCat,
		groupCat,
		groupCatLs,
		unwindCatGroup,
		replaceRoot,
	})

	if err != nil {
		return nil, err
	}

	for cursor.Next(c) {
		var rc domain.RoutineCategory

		if err := cursor.Decode(&rc); err != nil {
			panic(err)
		}

		rcs = append(rcs, rc)
	}

	return rcs, nil
}
