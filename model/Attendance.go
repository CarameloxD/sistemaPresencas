package model

import "gorm.io/gorm"

type Attendance struct {
	gorm.Model `swaggerignore:"true"`
	Id         int      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	IdStudent  int      `gorm:"not null" json:"IdStudent"`
	Student    Student  `gorm:"ForeignKey: IdStudent"`
	IdSchedule int      `gorm:"not null" json:"IdSchedule"`
	Schedule   Schedule `gorm:"ForeignKey: IdSchedule"`
}
