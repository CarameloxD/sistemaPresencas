package model

import "gorm.io/gorm"

type Classroom struct {
	gorm.Model `swaggerignore:"true"`
	Id         int    `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Identifier int    `gorm:"type:varchar(3);not null" json:"identifier"`
	Capacity   string `json:"type:varchar(4);not null" json:"capacity"`
}
