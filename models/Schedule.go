package models

type Schedule struct {
	Date      string  `json:"date"`
	Interest  float64 `json:"interest"`
	Principal float64 `json:"principal"`
}

func NewSchedule(date string, principal float64, interest float64) *Schedule {
	return &Schedule{Date: date, Principal: principal, Interest: interest}
}
