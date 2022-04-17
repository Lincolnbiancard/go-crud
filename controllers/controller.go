package controllers

import (
	"net/http"

	"github.com/Lincolnbiancard/gin-api/database"
	"github.com/Lincolnbiancard/gin-api/models"
	"github.com/gin-gonic/gin"
)

func StudentsList(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func Welcome(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API diz: ": "E aí " + name + ", tudo certo?",
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func FindStudentById(request *gin.Context) {
	var student models.Student
	id := request.Params.ByName("id")
	database.DB.First(&student, &id)
	if student.ID == 0 {
		request.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}
	request.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var Student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&Student, id)
	c.JSON(http.StatusOK, gin.H{"data": "The student has been deleted"})
}

func UpdateStudent(c *gin.Context) {
	var Student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&Student, id)

	if err := c.ShouldBindJSON(&Student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&Student).UpdateColumns(Student)
	c.JSON(http.StatusOK, Student)
}

func SearchStudent(c *gin.Context) {
	var Student models.Student
	cpf := c.Param("cpf")
	database.DB.Where(&models.Student{Cpf: cpf}).First(&Student)

	if Student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, Student)
}
