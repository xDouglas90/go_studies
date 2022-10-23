package controllers

import (
	"github.com/xdouglas90/go_studies/gin-rest-api/models"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Greeting(c *gin.Context) {
	name := c.Param("name")

	c.JSON(200, gin.H{
		"message": "Olá " + name + " seja muito bem vindo(a)!",
	})
}
