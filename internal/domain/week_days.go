package domain

type WeekDays struct {
	Name       string `bson:"name,omitempty" json:"name,omitempty"`
	NumericDay  int `bson:"numeric_day,omitempty" json:"numeric_day,omitempty"`
}
