package model

import "time"

type Schedule struct {
	Id           int       `gorm:"primaryKey;autoIncrement:true" json:"id"`
	IdClass      int       `gorm:"not null" json:"idClass"`
	Class        Class     `gorm:"ForeignKey: IdClass"`
	IdSubject    int       `gorm:"not null" json:"idSubject"`
	Subject      Subject   `gorm:"ForeignKey: IdSubject"`
	StartingTime time.Time `json:"startingTime"`
	EndingTime   time.Time `json:"endingTime"`
	IdTeacher    int       `gorm:"not null" json:"idTeacher"`
	Teacher      Teacher   `gorm:"ForeignKey: IdTeacher"`
	IdClassroom  int       `gorm:"not null" json:"idClassroom"`
	Classroom    Classroom `gorm:"ForeignKey: IdClassroom"`
}
