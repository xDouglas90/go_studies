package main

import (
	"net/http"

	"studies/alura/store/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
