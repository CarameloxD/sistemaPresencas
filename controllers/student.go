package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
	"time"
)

func GetStudentByNumber(c *gin.Context) {
	var user model.Student
	services.OpenDatabase()
	services.Db.Find(&user, c.Param("student_number"))

	if user.StudentNumber == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}

	var subjectsName string
	var schedulesStartingTime, schedulesEndingtime time.Time
	row := services.Db.Raw("Select subjects.name, schedules.starting_time, schedules.ending_time from schedules, classes, subjects, classrooms, teachers, students, subscriptions where subscriptions.id_student = students.id and classes.id = subscriptions.id_class and subjects.id = schedules.id_subject and teachers.id = schedules.id_teacher and classrooms.id = schedules.id_classroom and classes.id = schedules.id_class and students.student_number = ?", user.StudentNumber).Row()
	row.Scan(&subjectsName, &schedulesStartingTime, &schedulesEndingtime)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user, "subjectsName": subjectsName, "schedulesStartingTime": schedulesStartingTime, "schedulesEndingtime": schedulesEndingtime})
}

func InsertStudent(c *gin.Context) {
	//services.OpenDatabase()
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}

	services.Db.Save(&student)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})
}

func GetAllStudents(c *gin.Context) {
	var students []model.Student

	services.OpenDatabase()
	rows, _ := services.Db.Raw("Select id, name from students").Rows()

	for rows.Next() {
		services.Db.ScanRows(rows, &students)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "students": students})
}

func DeleteStudent(c *gin.Context) {
	var student model.Student
	id := c.Param("id")
	services.Db.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound,gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	} 
	services.Db.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeded!"})
}