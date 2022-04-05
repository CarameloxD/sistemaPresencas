package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model `swaggerignore:"true"`
	Id         int    `gorm:"primaryKey;autoIncrement:true" json:"Id"`
	Name       string `gorm:"size:255;not null" json:"name"`
	Username   string `gorm:"size:255;not null" json:"username"`
	Email      string `gorm:"size:255;not null" json:"email"`
	Picture    string `gorm:"size:255;not null" json:"picture"`
}
