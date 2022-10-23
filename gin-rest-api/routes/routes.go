package routes

import (
	"github.com/xdouglas90/go_studies/gin-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetStudents)
	r.GET("/:name", controllers.Greeting)
	r.Run()
}
