package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
)

func InsertSubject(c *gin.Context) {
	var subject model.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}
	services.Db.Save(&subject)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})
}

func GetAllSubjects(c *gin.Context) {
	var subjects []model.Subject

	services.OpenDatabase()
	services.Db.Select("*").Find(&subjects)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "subjects": subjects})
}

func DeleteSubject(c *gin.Context) {
	var subject model.Subject
	services.OpenDatabase()
	services.Db.Find(&subject, c.Param("id"))

	if subject.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "subject not found!"})
		return
	}
	services.Db.Delete(&subject)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Deleted!"})
}
