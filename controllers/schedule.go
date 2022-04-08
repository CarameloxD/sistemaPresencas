package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
	"strings"
	"time"
)

type Request struct {
	IdClass      int
	StartingTime string
	EndingTime   string
	IdClassroom  int
}

func InsertSchedule(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}
	request.StartingTime = strings.Replace(request.StartingTime, ".000", "", 1)
	request.EndingTime = strings.Replace(request.EndingTime, ".000", "", 1)

	const layout = "2006-01-02  15:04:05"
	var schedule model.Schedule
	schedule.IdClass = request.IdClass
	schedule.StartingTime, _ = time.Parse(layout, request.StartingTime)
	schedule.EndingTime, _ = time.Parse(layout, request.EndingTime)
	schedule.IdClassroom = request.IdClassroom
	services.Db.Save(&schedule)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})
}
