package handlers

import (
	"encoding/json"
	"hospital-portal/internal/database"
	"hospital-portal/internal/database/model"
	util "hospital-portal/internal/utils"
	"log"

	"net/http"
)

type LoginRequestBody struct{
	Username string `json:"username"`
	Password string	`json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	loginCreds := new(LoginRequestBody)
	err := json.NewDecoder(r.Body).Decode(loginCreds)
	if err != nil {
		http.Error(w, "couldnt decode request body", http.StatusInternalServerError)
		return
	}

	db := database.GetDB()

	var user model.User
	db.Where("Username = ?", loginCreds.Username).First(&user)
	if(util.VerifyPassword(loginCreds.Password,user.Password)){
		tokenString, err := util.CreateToken(user.Username,user.Role)
		if err != nil {
			log.Println("coudnlt generate token: ",err)
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}
		util.SetAuthCookie(w,tokenString)
		return
	}else{
		log.Println("error login : ",err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
}
