package model

type Class struct {
	Id           int     `gorm:"primaryKey;autoIncrement:true" json:"id"`
	ClassAcronym string  `gorm:"type:varchar(3);not null" json:"classAcronym"`
	IdSubject    int     `gorm:"not null" json:"idSubject"`
	Subject      Subject `gorm:"ForeignKey: IdSubject"`
	IdTeacher    int     `gorm:"not null" json:"idTeacher"`
	Teacher      Teacher `gorm:"ForeignKey: IdTeacher"`
}
