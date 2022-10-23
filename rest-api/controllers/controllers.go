package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.PoaPersonalities)
}
