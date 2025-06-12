package model

import "gorm.io/gorm"

type Patient struct {
	Name        string         `gorm:"unique; not null" json:"name"`
	Age         int            `gorm:"not null" json:"age"`
	Gender      string         `gorm:"not null" json:"gender"`
	Phone       string         `gorm:"not null" json:"phone"`
	Address     string         `gorm:"not null" json:"address"`
	BloodType 	string			`gorm:"not null" json:"bloodType"`
	Weight		int				`gorm:"not null" json:"weight"`
	Diagnosis     string		`gorm:"not null" json:"diagnosis"`
	ReceptionistId uint			`json:"receptionistId"`
	Receptionist   User   		`gorm:"foreignKey:ReceptionistId"`
	gorm.Model
}