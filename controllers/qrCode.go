package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
)

func GetQrCode(c *gin.Context) {
	var classroom model.Classroom
	services.OpenDatabase()
	services.Db.Where("id_classroom = ?", c.Param("id")).First(&classroom)

	if classroom.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Classroom not found!"})
		return
	}

	var schedules []model.Schedule
	rows, _ := services.Db.Raw("Select schedules.id from schedules, classrooms where classrooms.identifier = ? and schedules.id_classroom = classrooms.id and date(schedules.starting_time) = current_date order by schedules.starting_time", classroom.Identifier).Rows()
	for rows.Next() {
		services.Db.ScanRows(rows, &schedules)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "schedules": schedules})
}
