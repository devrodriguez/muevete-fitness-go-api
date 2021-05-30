package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoutineSchedule struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Routine  Routine            `bson:"routine,omitempty" json:"routine,omitempty"`
	WeekDays WeekDays           `bson:"week_days,omitempty" json:"week_days,omitempty"`
}

type RoutineScheduleMod struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Routine  primitive.ObjectID `bson:"routine,omitempty" json:"routine,omitempty"`
	WeekDays primitive.ObjectID `bson:"week_days,omitempty" json:"week_days,omitempty"`
}
