package domain

type RoutineSchedule struct {
	Routine Routine `bson:"routine,omitempty" json:"routine,omitempty"`
	WeekDays WeekDays `bson:"week_days,omitempty" json:"week_days,omitempty"`
}
