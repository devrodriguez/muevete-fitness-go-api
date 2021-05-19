package domain

type Schedule struct {
	Weekly `bson:"weekly,omitempty" json:"weekly,omitempty"`
	Customer `bson:"customer,omitempty" json:"customer,omitempty"`
}
