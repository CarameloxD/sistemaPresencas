package model

import (
	"gorm.io/gorm"
	"time"
)

type Schedule struct {
	gorm.Model   `swaggerignore:"true"`
	Id           int       `gorm:"primaryKey;autoIncrement:true" json:"id"`
	IdClass      int       `gorm:"not null" json:"idClass"`
	Class        Class     `gorm:"ForeignKey: IdClass"`
	StartingTime time.Time `json:"startingTime"`
	EndingTime   time.Time `json:"endingTime"`
	IdClassroom  int       `gorm:"not null" json:"idClassroom"`
	Classroom    Classroom `gorm:"ForeignKey: IdClassroom"`
}
