package main

import (
	"rest-api/database"
	"rest-api/routes"
)

func main() {
	database.Connect()

	routes.HandleRequests()
}
