package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoutineCategory struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Routines []Routine          `bson:"routines,omitempty" json:"routine,omitempty"`
	Category Category           `bson:"category,omitempty" json:"category,omitempty"`
}
