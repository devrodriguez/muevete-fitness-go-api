package domain

type Session struct {
	Name      string `bson:"name,omitempty" json:"name,omitempty"`
	StartHour string `bson:"start_hour,omitempty" json:"start_hour,omitempty"`
	FinalHour string `bson:"final_hour,omitempty" json:"final_hour,omitempty"`
	Period    string `bson:"period,omitempty" json:"period,omitempty"`
}
