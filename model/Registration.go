package model

type Registration struct {
	Id        int     `gorm:"primaryKey;autoIncrement:true" json:"id"`
	IdClass   int     `gorm:"not null" json:"idClass"`
	Class     Class   `gorm:"ForeignKey: IdClass"`
	IdStudent int     `gorm:"not null" json:"idStudent"`
	Student   Student `gorm:"ForeignKey: IdStudent"`
}
