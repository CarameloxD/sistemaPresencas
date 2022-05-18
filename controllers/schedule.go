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

type RequestSC struct {
	Id           int
	IdClass      int
	StartingTime time.Time
	EndingTime   time.Time
	IdClassroom  int
	Name         string
}

type SC struct {
	IdClass     int
	IdSchedules []int
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

func GetSchedulesByClass(c *gin.Context) {
	var schedules []RequestSC
	var class model.Class
	services.Db.Find(&class, c.Param("id"))

	services.OpenDatabase()
	rows, _ := services.Db.Raw("Select distinct schedules.*, subjects.name from schedules, classes, subjects where schedules.id_class = ? and subjects.id = classes.id_subject and  schedules.id_class = classes.id and schedules.deleted_at is null", class.Id).Rows()

	for rows.Next() {
		services.Db.ScanRows(rows, &schedules)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "schedules": schedules})
}

func DeleteSchedule(c *gin.Context) {
	var schedule SC
	services.OpenDatabase()
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}

	for _, idSchedule := range schedule.IdSchedules {
		var schedules2 []model.Schedule

		rows, _ := services.Db.Raw("Select * from schedules where id_class = ? and id = ?", schedule.IdClass, idSchedule).Rows()
		for rows.Next() {
			services.Db.ScanRows(rows, &schedules2)
			services.Db.Delete(&schedules2)
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Deleted!"})
}
