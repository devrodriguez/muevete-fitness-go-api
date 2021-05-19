package domain

type Customer struct {
	Name     string `bson:"name,omitempty" json:"name,omitempty"`
	LastName string `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
}
