package routes

import (
	"log"
	"net/http"
	"rest-api/controllers"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.SinglePersonality).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("PUT")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonalityName).Methods("PATCH")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonalityHistory).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8080", r))
}
