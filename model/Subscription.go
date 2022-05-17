package model

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model `swaggerignore:"true"`
	Id         int     `gorm:"primaryKey;autoIncrement:true" json:"id"`
	IdCourse   int     `gorm:"not null" json:"idCourse"`
	Course     Course  `gorm:"ForeignKey: IdCourse"`
	IdStudent  int     `gorm:"not null" json:"idStudent"`
	Student    Student `gorm:"ForeignKey: IdStudent"`
}
