package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
)

func GetAllCourses(c *gin.Context) {
	var courses []model.Course

	services.OpenDatabase()
	rows, _ := services.Db.Raw("Select * from courses").Rows()

	for rows.Next() {
		services.Db.ScanRows(rows, &courses)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "courses": courses})
}

func InsertCourse(c *gin.Context) {
	var course model.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}
	services.Db.Save(&course)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})
}
