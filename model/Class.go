package model

import "time"

type Class struct {
	Subject      string    `json:"subject"`
	Classroom    string    `json:"classroom"`
	Teacher      string    `json:"teacher"`
	ClassAcronym string    `json:"classAcronym"`
	Date         time.Time `json:"date"`
	StartingTime time.Time `json:"startingTime"`
	EndingTime   time.Time `json:"endingTime"`
}
