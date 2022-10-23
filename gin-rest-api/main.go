package main

import (
	"github.com/xdouglas90/go_studies/gin-rest-api/database"
	"github.com/xdouglas90/go_studies/gin-rest-api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
