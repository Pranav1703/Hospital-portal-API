package handlers

import (
	"encoding/json"
	"fmt"
	"hospital-portal/internal/database"
	"hospital-portal/internal/database/model"
	util "hospital-portal/internal/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterPatient(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserContextKey).(AuthUser)
	if user.Role != "receptionist" {
		http.Error(w, "Only receptionists can register patients", http.StatusForbidden)
		return
	}

	var patient model.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	patient.ReceptionistId = util.GetUserIDByUsername(user.Username)
	if err := database.GetDB().Create(&patient).Error; err != nil {
		http.Error(w, "Failed to save patient", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Patient record added. Patient ID: %d",patient.ID)))
}

func GetAllPatients(w http.ResponseWriter, r *http.Request) {
	var patients []model.Patient
	if err := database.GetDB().Preload("Receptionist").Find(&patients).Error; err != nil {
		http.Error(w, "Failed to retrieve patients", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(patients)
}


func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var patient model.Patient

	db := database.GetDB()
	if err := db.First(&patient, id).Error; err != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if err := db.Save(&patient).Error; err != nil {
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Patient record updated."))
}

func DeletePatient(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserContextKey).(AuthUser)
	if user.Role != "receptionist" {
		http.Error(w, "Only receptionists can delete patients", http.StatusForbidden)
		return
	}

	id := chi.URLParam(r, "id")
	if err := database.GetDB().Unscoped().Delete(&model.Patient{}, id).Error; err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Patient record deleted."))
}
