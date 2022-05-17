package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
)

func InsertClassroom(c *gin.Context) {
	var classroom model.Classroom
	if err := c.ShouldBindJSON(&classroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}
	services.Db.Save(&classroom)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})
}

func GetAllClassrooms(c *gin.Context) {
	var classrooms []model.Classroom

	services.OpenDatabase()
	services.Db.Select("id,identifier,capacity").Find(&classrooms)

	if len(classrooms) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "classrooms not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "classrooms": classrooms})
}

func DeleteClassroom(c *gin.Context) {
	var classroom model.Classroom
	services.OpenDatabase()
	services.Db.Find(&classroom, c.Param("id"))

	if classroom.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Classroom not found!"})
		return
	}
	services.Db.Delete(&classroom)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Classroom deleted!"})
}
