package model

type Class struct {
	Id           int    `gorm:"primaryKey;autoIncrement:true" json:"id"`
	ClassAcronym string `gorm:"type:varchar(3);not null" json:"classAcronym"`
}
