package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	LastName string             `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
}
