package model

type Course struct {
	Id    int    `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Title string `gorm:"size:255;not null" json:"title"`
}
