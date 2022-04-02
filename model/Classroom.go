package model

import "gorm.io/gorm"

type Classroom struct {
	gorm.Model `swaggerignore:"true"`
	Id         int `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Identifier int `json:"identifier"`
	Capacity   int `json:"capacity"`
}
