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

func GetStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")
	database.DBConn.Where("cpf = ?", cpf).First(&student)
	if student.CPF == cpf {
		c.JSON(200, student)
	} else {
		c.JSON(404, gin.H{"message": "Aluno não encontrado"})
	}
}

func GetStudentByRG(c *gin.Context) {
	var student models.Student
	rg := c.Param("rg")
	database.DBConn.Where("rg = ?", rg).First(&student)
	if student.RG == rg {
		c.JSON(200, student)
	} else {
		c.JSON(404, gin.H{"message": "Aluno não encontrado"})
	}
}
