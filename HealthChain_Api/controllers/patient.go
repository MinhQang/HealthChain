package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"HealthChain_API/config"
	"HealthChain_API/models"
	"github.com/gorilla/mux"
)

func CreatePatientsController(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&patient).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(patient)
}

func GetPatientsController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	var patients models.Patient
	if err := config.DB.Preload("EmergencyContact").First(&patients, id).Error; err != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(patients)
}

func UpdatePatientsController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	var patient models.Patient
	if err := config.DB.First(&patient, id).Error; err != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	config.DB.Save(&patient)
	json.NewEncoder(w).Encode(patient)
}

func DeletePatientsController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	if err := config.DB.Delete(&models.Patient{}, id).Error; err != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
