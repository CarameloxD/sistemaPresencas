package model

import "gorm.io/gorm"

//Cadeira
type Subject struct {
	gorm.Model `swaggerignore:"true"`
	Id         int    `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name       string `gorm:"size:255;not null" json:"name"`
	Type       string `gorm:"size:255;not null" json:"type"`
}
