package model

import "time"

type Schedule struct {
	IdClass int `json:"idClass"`
	IdSubject int `json:"idSubject"`
	StartingTime time.Time `json:"startingTime"`
	EndingTime time.Time `json:"endingTime"`
	IdTeacher int `json:"idTeacher"`
	IdClassroom int `json:"idClassroom"`
}
