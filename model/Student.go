package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model `swaggerignore:"true"`
	StudentNumber int `json:"studentnumber"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Picture    string `json:"picture"`
}
