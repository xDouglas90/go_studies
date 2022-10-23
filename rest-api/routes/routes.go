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
	log.Fatal(http.ListenAndServe(":8080", r))
}
