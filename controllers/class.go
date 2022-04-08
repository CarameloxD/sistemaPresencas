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
	rows, _ := services.Db.Raw("Select * from classes").Rows()

	for rows.Next() {
		services.Db.ScanRows(rows, &classes)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "classes": classes})
}
