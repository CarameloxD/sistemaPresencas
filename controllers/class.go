package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
)

func InsertClass(c *gin.Context) {
	var class model.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}
	services.Db.Save(&class)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})
}

func GetAllClasses(c *gin.Context) {
	var classes []model.Class

	services.OpenDatabase()
	services.Db.Select("id,class_acronym,id_subject,id_teacher").Find(&classes)
	if len(classes) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "classes not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "classes": classes})
}

func DeleteClass(c *gin.Context) {
	var class model.Class
	services.OpenDatabase()
	services.Db.Find(&class, c.Param("id"))

	if class.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Class not found!"})
		return
	}
	services.Db.Delete(&class)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Class deleted!"})
}
