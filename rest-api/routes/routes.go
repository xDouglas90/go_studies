package routes

import (
	"log"
	"net/http"
	"studies/alura/rest-api/controllers"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
