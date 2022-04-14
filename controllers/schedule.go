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

type RequestGet struct {
	Id           int
	StartingTime string
	ClassAcronym string
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

func GetAllSchedules(c *gin.Context) {
	var schedules []RequestGet

	services.OpenDatabase()
	rows, _ := services.Db.Raw("select schedules.id, schedules.starting_time, classes.class_acronym from schedules, classes where schedules.id_class = classes.id;").Rows()

	for rows.Next() {
		services.Db.ScanRows(rows, &schedules)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "schedules": schedules})
}

func GetStudentsBySchedule(c *gin.Context) {
	var schedule model.Schedule
	services.OpenDatabase()
	services.Db.Find(&schedule, c.Param("id"))

	if schedule.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Schedule not found!"})
		return
	}

	var students []model.Student
	rows, _ := services.Db.Raw("Select students.* from schedules, classes, subjects, courses, subscriptions, students where schedules.id = ? and schedules.id_class = classes.id and classes.id_subject = subjects.id and subjects.id_course = courses.id and subscriptions.id_course = courses.id and subscriptions.id_student = students.id", schedule.Id).Rows()
	for rows.Next() {
		services.Db.ScanRows(rows, &students)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": schedule, "students": students})
}
