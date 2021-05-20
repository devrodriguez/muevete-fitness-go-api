package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type WeekDays struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string `bson:"name,omitempty" json:"name,omitempty"`
	NumericDay int    `bson:"numeric_day,omitempty" json:"numeric_day,omitempty"`
}
