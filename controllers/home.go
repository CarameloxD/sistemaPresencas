package controllers

import (
	"fmt"
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
	row := services.Db.Raw("Select subjects.name, schedules.starting_time, schedules.ending_time from schedules, classes, subjects, classrooms, teachers, students, registrations where registrations.id_student = students.id and classes.id = registrations.id_class and subjects.id = schedules.id_subject and teachers.id = schedules.id_teacher and classrooms.id = schedules.id_classroom and classes.id = schedules.id_class and students.student_number = ?", user.StudentNumber).Row()
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

	fmt.Print(&student)
	//fmt.Print(student.Name, student.Email)
	services.Db.Save(&student)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})

	//c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully", "student_number": student.StudentNumber})

	//c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": c.Param("email")})
}

func InsertClassroom(c *gin.Context) {
	var classroom model.Classroom
	if err := c.ShouldBindJSON(&classroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}
	services.Db.Save(&classroom)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})
}

func InsertClass(c *gin.Context) {
	var class model.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}
	services.Db.Save(&class)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})
}
