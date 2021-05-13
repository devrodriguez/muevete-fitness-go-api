package domain

type Routine struct {
	Name        string `bson:"name,omitempty" json:"name,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}
