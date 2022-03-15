package model

import "gorm.io/gorm"

type Classroom struct {
	gorm.Model `swaggerignore:"true"`
	Identifier int `json:"identifier"`
	Capacity       string `json:"capacity"`
}