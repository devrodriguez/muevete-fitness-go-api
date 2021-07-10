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
