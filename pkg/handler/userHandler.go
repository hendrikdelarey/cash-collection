package handler

import (
	"encoding/json"
	"net/http"
)

type UserRegistrationDTO struct {
	Token        string `json:"token"`
	ClientID     string `json:"investecClientId"`
	ClientSecret string `json:"investecClientSecret"`
}

type UserLoginDTO struct {
	Token        string `json:"token"`
}

// TODO: correct dto required fields
func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dto UserRegistrationDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(dto)
}

// TODO: return session token
func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dto UserLoginDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(dto)
}
