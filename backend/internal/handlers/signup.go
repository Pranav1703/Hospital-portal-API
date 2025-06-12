package handlers

import (
	"encoding/json"
	"hospital-portal/internal/database"
	"hospital-portal/internal/database/model"
	"log"
	"net/http"
)

type ReqBody struct {
	Username string
	Password string
	Role     string
}

func Signup(w http.ResponseWriter, r *http.Request) {
	loginCreds := new(ReqBody)
	err := json.NewDecoder(r.Body).Decode(loginCreds)
	if err != nil {
		http.Error(w, "couldnt decode request body", http.StatusInternalServerError)
		return
	}

	newUser := &model.User{
		Username: loginCreds.Username,
		Password: loginCreds.Password,
		Role: loginCreds.Role,
	}

	db := database.GetDB()
	result := db.Create(&newUser)
	if result.Error != nil {
		log.Println(result.Error)
		http.Error(w,result.Error.Error(),500)
	}

}