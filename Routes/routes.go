package Routes

import (
	"github.com/EversonRubira/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Salutation)
	r.POST("/students", controllers.CreateNewStudent)
	r.GET("/student/:id", controllers.SearchStudentById)
	r.DELETE("/students/: id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCPF)
	r.Run()

}
