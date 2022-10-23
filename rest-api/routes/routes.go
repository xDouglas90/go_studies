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
	r.HandleFunc("/api/personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8080", r))
}
