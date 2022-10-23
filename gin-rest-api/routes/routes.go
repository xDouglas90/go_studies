package routes

import (
	"github.com/xdouglas90/go_studies/gin-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetStudents)
	r.GET("/students/:id", controllers.GetStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.GET("/:name", controllers.Greeting)
	r.POST("/students", controllers.CreateStudent)
	r.Run()
}
