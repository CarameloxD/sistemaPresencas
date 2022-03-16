package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model    `swaggerignore:"true"`
	Id            int    `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name          string `gorm:"size:255;not null" json:"name"`
	Email         string `gorm:"size:255;not null" json:"email"`
	StudentNumber int    `json:"studentnumber"`
	Picture       string `json:"picture"`
}
