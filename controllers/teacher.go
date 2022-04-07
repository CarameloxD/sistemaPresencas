package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
	//"time"
)

func InsertTeacher(c *gin.Context) {
	//services.OpenDatabase()
	var teacher model.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}

	fmt.Print(&teacher)
	//fmt.Print(student.Name, student.Email)
	services.Db.Save(&teacher)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})

	//c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully", "student_number": student.StudentNumber})

	//c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": c.Param("email")})
}

func GetTeacherInfo(c *gin.Context) {
		var user model.Teacher
		services.OpenDatabase()
		services.Db.Find(&user, c.Param("username"))
	
		if len(user.Username) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
			return
		}
	
		//var subjectsName string
		//var schedulesStartingTime, schedulesEndingtime time.Time
		//row := services.Db.Raw("Select subjects.name, schedules.starting_time, schedules.ending_time from schedules, classes, subjects, classrooms, teachers, students, registrations where registrations.id_student = students.id and classes.id = registrations.id_class and subjects.id = schedules.id_subject and teachers.id = schedules.id_teacher and classrooms.id = schedules.id_classroom and classes.id = schedules.id_class and schedules.id_teacher = ?", user.ID).Row()
		//row.Scan(&subjectsName, &schedulesStartingTime, &schedulesEndingtime)
		//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user, "subjectsName": subjectsName, "schedulesStartingTime": schedulesStartingTime, "schedulesEndingtime": schedulesEndingtime})
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

func DeleteTeacher(c *gin.Context) {
	var teacher model.Teacher
	username := c.Param("username")
	services.Db.First(&teacher, username)
	if teacher.ID == 0 {
		c.JSON(http.StatusNotFound,gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	} 
	services.Db.Delete(&teacher)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeded!"})
}