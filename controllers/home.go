package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
)

func GetStudentByNumber(c *gin.Context) {
	var user model.Student
	services.OpenDatabase()
	services.Db.Find(&user, c.Param("student_number"))

	/*results := services.Db.Table("schedules, classes, subjects, classrooms, teachers, students, registrations").
	Select("subjects.name, schedules.starting_time, schedules.ending_time").
	Joins("where registrations.id_student = students.id and classes.id = registrations.id_class and subjects.id = schedules.id_subject and teachers.id = schedules.id_teacher and classrooms.id = schedules.id_classroom and classes.id = schedules.id_class")*/
	if user.StudentNumber == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}
