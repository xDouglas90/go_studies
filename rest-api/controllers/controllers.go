package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/database"
	"rest-api/models"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality

	database.DB.Find(&personalities)

	json.NewEncoder(w).Encode(personalities)
}

func SinglePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var personality models.Personality

	database.DB.First(&personality, key)

	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality

	json.NewDecoder(r.Body).Decode(&newPersonality)

	database.DB.Create(&newPersonality)

	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var personality models.Personality

	database.DB.First(&personality, key)
	database.DB.Delete(&personality)

	json.NewEncoder(w).Encode(personality)
}
