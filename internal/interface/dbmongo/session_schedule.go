package dbmongo

import (
	"context"
	"fmt"

	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDbSessionSchedule interface {
	GetAllSessionSchedule(context.Context) ([]domain.SessionSchedule, error)
	GetByWeekly(context.Context, string) (int64, error)
	SaveSessionSchedule(context.Context, domain.SessionScheduleMod) error
}

type ImpDbSessionSchedule struct {
	*mongo.Client
}

func NewDbSessionSchedule(cli *mongo.Client) IDbSessionSchedule {
	return &ImpDbSessionSchedule{
		cli,
	}
}

func (ss *ImpDbSessionSchedule) GetAllSessionSchedule(ctx context.Context) ([]domain.SessionSchedule, error) {
	var sess []domain.SessionSchedule

	docRef := ss.Client.Database("fitness").Collection("session_schedule")
	lookCust := bson.D{
		{"$lookup", bson.D{
			{"from", "customers"},
			{"localField", "customer"},
			{"foreignField", "_id"},
			{"as", "customer"},
		}},
	}

	unwindCust := bson.D{
		{"$unwind", bson.D{
			{"path", "$customer"},
			{"preserveNullAndEmptyArrays", false},
		}},
	}

	lookWeek := bson.D{{"$lookup", bson.D{
		{"from", "weeklies"},
		{"localField", "weekly"},
		{"foreignField", "_id"},
		{"as", "weekly"},
	}}}

	unwindWeek := bson.D{
		{"$unwind", bson.D{
			{"path", "$weekly"},
			{"preserveNullAndEmptyArrays", false},
		}},
	}

	lookSes := bson.D{
		{"$lookup", bson.D{
			{"from", "sessions"},
			{"localField", "weekly.session"},
			{"foreignField", "_id"},
			{"as", "weekly.session"},
		}},
	}

	unwindSes := bson.D{
		{"$unwind", bson.D{
			{"path", "$weekly.session"},
			{"preserveNullAndEmptyArrays", false},
		}}}

	lookRouSch := bson.D{
		{"$lookup", bson.D{
			{"from", "routine_schedule"},
			{"localField", "weekly.routine_schedule"},
			{"foreignField", "_id"},
			{"as", "weekly.routine_schedule"},
		}},
	}

	unwindRouSch := bson.D{
		{"$unwind", bson.D{
			{"path", "$weekly.routine_schedule"},
			{"preserveNullAndEmptyArrays", false},
		}}}

	lookRouSchRou := bson.D{
		{"$lookup", bson.D{
			{"from", "routines"},
			{"localField", "weekly.routine_schedule.routine"},
			{"foreignField", "_id"},
			{"as", "weekly.routine_schedule.routine"},
		}}}

	unwindRouSchRou := bson.D{
		{"$unwind", bson.D{
			{"path", "$weekly.routine_schedule.routine"},
			{"preserveNullAndEmptyArrays", false},
		}}}

	lookRouSchWkd := bson.D{
		{"$lookup", bson.D{
			{"from", "week_days"},
			{"localField", "weekly.routine_schedule.week_day"},
			{"foreignField", "_id"},
			{"as", "weekly.routine_schedule.week_day"},
		}}}

	unwindRouSchWkd := bson.D{
		{"$unwind", bson.D{
			{"path", "$weekly.routine_schedule.week_day"},
			{"preserveNullAndEmptyArrays", false},
		}}}

	cursor, err := docRef.Aggregate(ctx, mongo.Pipeline{
		lookCust,
		unwindCust,
		lookWeek,
		unwindWeek,
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
		var ses domain.SessionSchedule

		if err := cursor.Decode(&ses); err != nil {
			panic(err)
		}

		sess = append(sess, ses)
	}

	return sess, nil
}

func (ss *ImpDbSessionSchedule) GetByWeekly(ctx context.Context, wkID string) (int64, error) {
	docID, err := primitive.ObjectIDFromHex(wkID)
	collRef := ss.Client.Database("fitness").Collection("session_schedule")
	filter := bson.M{"weekly": docID}
	nDocs, err := collRef.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return nDocs, nil
}

func (ss *ImpDbSessionSchedule) SaveSessionSchedule(ctx context.Context, sess domain.SessionScheduleMod) error {

	weeklyID, _ := primitive.ObjectIDFromHex(sess.WeeklyID.Hex())
	custID, _ := primitive.ObjectIDFromHex(sess.CustomerID.Hex())
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"weekly": weeklyID, "customer": custID}
	data := bson.M{"$set": sess}
	collRef := ss.Client.Database("fitness").Collection("session_schedule")

	res, err := collRef.UpdateOne(ctx, filter, data, opts)

	fmt.Printf("[updated:%d]\n", res.ModifiedCount)
	fmt.Printf("[upserted:%d]\n", res.UpsertedCount)
	fmt.Printf("[upserted_id:%v]\n", res.UpsertedID)

	if err != nil {
		return err
	}

	return nil
}
