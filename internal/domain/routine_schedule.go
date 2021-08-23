package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoutineSchedule struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Routine Routine            `bson:"routine,omitempty" json:"routine,omitempty"`
	WeekDay WeekDay            `bson:"week_day,omitempty" json:"week_day,omitempty"`
}

type RoutineScheduleMod struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Routine primitive.ObjectID `bson:"routine,omitempty" json:"routine,omitempty"`
	WeekDay primitive.ObjectID `bson:"week_day,omitempty" json:"week_day,omitempty"`
}
