package routes

import (
	"github.com/EversonRubira/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/Students", controllers.displaysAllStudents)
	r.GET("/:name", controllers.Salutation)
	r.POST("/students", controllers.createNewStudent)
	r.GET("/student/:id", controllers.searchStudentById)
	r.DELETE("/students/:id", controllers.deleteStudent)
	r.PATCH("/students/:id", controllers.editStudent)
	r.Run()

}
