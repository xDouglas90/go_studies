package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xdouglas90/go_studies/gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetStudents)
	r.GET("/students/:id", controllers.GetStudent)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)
	r.GET("/students/rg/:rg", controllers.GetStudentByRG)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.POST("/students", controllers.CreateStudent)
	r.Run()
}
