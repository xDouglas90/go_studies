package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xdouglas90/go_studies/gin-rest-api/database"
	"github.com/xdouglas90/go_studies/gin-rest-api/models"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DBConn.Find(&students)
	c.JSON(200, students)
}

func GetStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	database.DBConn.First(&student, id)
	if student.ID == 0 {
		c.JSON(404, gin.H{"error": "Aluno não encontrado"})
		return
	}
	c.JSON(200, student)
}

func Greeting(c *gin.Context) {
	name := c.Param("name")
	var student models.Student
	database.DBConn.Where("name = ?", name).First(&student)
	// verify if student name contains the name passed in the URL
	if student.Name == name {
		c.JSON(200, gin.H{"message": "Olá " + name + " seja muito bem vindo(a)!"})
	} else {
		c.JSON(404, gin.H{"message": "Aluno não encontrado"})
	}
}

func CreateStudent(c *gin.Context) {
	var newStudent models.Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DBConn.Create(&newStudent)
	c.JSON(201, newStudent)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	database.DBConn.First(&student, id)
	if student.ID == 0 {
		c.JSON(404, gin.H{"error": "Aluno não encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DBConn.Save(&student)
	c.JSON(200, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	database.DBConn.First(&student, id)
	if student.ID == 0 {
		c.JSON(404, gin.H{"error": "Aluno não encontrado"})
		return
	}
	database.DBConn.Delete(&student)
	c.JSON(204, nil)
}
