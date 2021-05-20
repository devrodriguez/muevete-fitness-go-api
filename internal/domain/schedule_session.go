package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ScheduleSession struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Weekly   `bson:"weekly,omitempty" json:"weekly,omitempty"`
	Customer `bson:"customer,omitempty" json:"customer,omitempty"`
}
