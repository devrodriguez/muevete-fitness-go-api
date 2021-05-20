package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoutineSchedule struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Routine  Routine            `bson:"routine,omitempty" json:"routine,omitempty"`
	WeekDays WeekDays           `bson:"week_days,omitempty" json:"week_days,omitempty"`
}
