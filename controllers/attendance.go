package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
)

type RequestA struct {
	IdSchedule int
	IdStudents []int
}

func InsertAttendance(c *gin.Context) {
	var request RequestA
	services.OpenDatabase()
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}

	for _, idStudent := range request.IdStudents {
		var attendance model.Attendance
		attendance.IdSchedule = request.IdSchedule
		attendance.IdStudent = idStudent
		services.Db.Save(&attendance)
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})
}
