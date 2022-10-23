package main

import (
	"github.com/xdouglas90/go_studies/gin-rest-api/database"
	"github.com/xdouglas90/go_studies/gin-rest-api/models"
	"github.com/xdouglas90/go_studies/gin-rest-api/routes"
)

func main() {
	database.Connect()
	models.Students = []models.Student{
		{Name: "John Doe", CPF: "123456789", RG: "987654321"},
		{Name: "Jane Doe", CPF: "987654321", RG: "123456789"},
	}
	routes.HandleRequests()
}
