package handlers

import (
	"encoding/json"
	"fmt"

	"hospital-portal/internal/database"
	"hospital-portal/internal/database/model"
	util "hospital-portal/internal/utils"
	"log"
	"net/http"
)

type SignUpReqBody struct {
	Username string
	Password string
	Role     string
}

func Signup(w http.ResponseWriter, r *http.Request) {
	details := new(SignUpReqBody)
	err := json.NewDecoder(r.Body).Decode(details)
	if err != nil {
		http.Error(w, "couldnt decode request body", http.StatusInternalServerError)
		return
	}

	hashedPass,err := util.HashPassword(details.Password)

	if err!= nil {
		fmt.Println("couldn't encrpt the password",err)
	}

	newUser := &model.User{
		Username: details.Username,
		Password: hashedPass,
		Role:     details.Role,
	}

	db := database.GetDB()
	result := db.Create(&newUser)
	if result.Error != nil {
		fmt.Println("Db error:",result.Error)
		http.Error(w, result.Error.Error(), 500)
	}
	log.Println("new user created. ID:",newUser.ID)
	w.Write([]byte(fmt.Sprintf("new user created. ID: %d",newUser.ID)))
}
