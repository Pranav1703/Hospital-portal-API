package model

import "gorm.io/gorm"

type User struct {
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"not null;check:role IN ('doctor','receptionist')" json:"role"`
	gorm.Model
}