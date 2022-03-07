package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Type       string `json:"type"`
	Picture    string `json:"picture"`
	//Agenda     Schedule `json:"agenda"`
}
