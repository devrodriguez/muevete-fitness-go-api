package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type SessionSchedule struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Weekly   `bson:"weekly,omitempty" json:"weekly,omitempty"`
	Customer `bson:"customer,omitempty" json:"customer,omitempty"`
}

type SessionScheduleMod struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	WeeklyID   primitive.ObjectID `bson:"weekly,omitempty" json:"weekly,omitempty"`
	CustomerID primitive.ObjectID `bson:"customer,omitempty" json:"customer,omitempty"`
}
