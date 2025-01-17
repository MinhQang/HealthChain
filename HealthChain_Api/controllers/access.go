package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"HealthChain_API/models"
	"github.com/gorilla/mux"
)

func AddAccess(w http.ResponseWriter, r *http.Request) {
	var access models.Access
	err := json.NewDecoder(r.Body).Decode(&access)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = models.CreateAccess(&access)
	if err != nil {
		http.Error(w, "Failed to create access", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(access)
}

func GetAccess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["userID"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	var access []models.Access
	err = models.GetAccessByID(userID, &access)
	if err != nil {
		http.Error(w, "Access not found", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(access)
}

func DeleteAccess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accessID, err := strconv.ParseUint(vars["accessID"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	var access models.Access
	access.ID = accessID
	err = models.DeleteAccess(&access)
	if err != nil {
		http.Error(w, "Failed to delete access", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
