package handlers

import (
	"encoding/json"
	"hospital-portal/internal/database"
	
	"net/http"
)


func Login(w http.ResponseWriter, r *http.Request) {
	loginCreds := new(ReqBody)
	err := json.NewDecoder(r.Body).Decode(loginCreds)
	if err != nil {
		http.Error(w,"couldnt decode request body",http.StatusInternalServerError)
		return 
	}
	_ = database.GetDB()


}