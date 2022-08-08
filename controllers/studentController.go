package controllers

import (
	"net/http"

	"github.com/EversonRubira/api-go-gin/database"
	"github.com/EversonRubira/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func ShowAllStudents(c *gin.Context) {
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

func CreateNewStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})

	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func SearchStudentById(c *gin.Context) {
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

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})

}

func EditStudent(c *gin.Context) {
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

func SearchStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno Nao Encontrado"})
		return
	}
	c.JSON(http.StatusOK, student)

}
