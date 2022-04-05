package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model `swaggerignore:"true"`
	Id         int    `gorm:"primaryKey;autoIncrement:true"`
	Name       string `gorm:"size:255;not null" json:"name"`
	Username   string `gorm:"size:255;not null" json:"username"`
	Password   string `json:"password"`
	Email      string `gorm:"size:255;" json:"email"`
	Picture    string `gorm:"size:255;" json:"picture"`
}
