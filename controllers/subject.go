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
