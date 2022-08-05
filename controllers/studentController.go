package controllers

import (
	"net/http"

	"github.com/EversonRubira/api-go-gin/database"
	"github.com/EversonRubira/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func displaysAllStudents(c *gin.Context) {
	var student []models.Student
	database.DB.Find(&student)
	c.JSON(200, student)
}

func Salutation(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + name + ", tudo beleza?",
	})

}

func createNewStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})

	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func searchStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno Nao Encontrado"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func deleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})

}

func editStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)

}
