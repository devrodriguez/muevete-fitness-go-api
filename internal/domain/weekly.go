package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Weekly struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Session         `bson:"session,omitempty" json:"session,omitempty"`
	RoutineSchedule `bson:"routine_schedule,omitempty" json:"routine_schedule,omitempty"`
	Status          bool `bson:"status,omitempty" json:"status,omitempty"`
}

type WeeklyMod struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Session         primitive.ObjectID `bson:"session,omitempty" json:"session,omitempty"`
	RoutineSchedule primitive.ObjectID `bson:"routine_schedule,omitempty" json:"routine_schedule,omitempty"`
	Status          bool `bson:"status,omitempty" json:"status,omitempty"`
}
