package domain

type Category struct {
	Name string `bson:"name,omitempty" json:"name,omitempty"`
}
