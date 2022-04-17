package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Lincolnbiancard/gin-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.StudentsList)
	r.GET("/students/:id", controllers.FindStudentById)
	r.POST("/students", controllers.CreateStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.GET("/students/cpf/:cpf", controllers.SearchStudent)
	r.Run()
}