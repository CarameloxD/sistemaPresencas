package model

import "gorm.io/gorm"

//Cadeira
type Subject struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name"`
	Type      string `json:"type"`
}