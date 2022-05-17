package model

import "gorm.io/gorm"

type Course struct {
	gorm.Model `swaggerignore:"true"`
	Id         int    `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Title      string `gorm:"size:255;not null" json:"title"`
}
