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
	rows, _ := services.Db.Raw("Select * from classrooms").Rows()

	for rows.Next() {
		services.Db.ScanRows(rows, &classrooms)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "classrooms": classrooms})
}
