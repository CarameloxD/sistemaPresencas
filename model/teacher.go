package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Picture    string `json:"picture"`
}