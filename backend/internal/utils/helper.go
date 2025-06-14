package util

import (
	"hospital-portal/internal/database"
	"hospital-portal/internal/database/model"
)

func GetUserIDByUsername(username string) uint {
	db := database.GetDB()
	var user model.User
	db.Where("username = ?", username).First(&user)
	return user.ID
}
