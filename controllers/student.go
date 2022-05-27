package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
	"time"
)

type RequestS struct {
	Id           int
	IdClass      int
	StartingTime time.Time
	EndingTime   time.Time
	IdClassroom  int
	Name         string
	Type         string
	Identifier   int
	Teacher      string
}

func GetStudentByNumber(c *gin.Context) {
	var user model.Student
	services.OpenDatabase()
	services.Db.Where("student_number = ?", c.Param("id")).First(&user)

	if user.StudentNumber == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "user": user})
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
	services.Db.Select("id,name,email,student_number").Find(&students)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "students": students})
}

func DeleteStudent(c *gin.Context) {
	var user model.Student
	services.OpenDatabase()
	services.Db.Find(&user, c.Param("id"))

	if user.StudentNumber == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}
	services.Db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Deleted!"})
}

func GetSchedulesByStudent(c *gin.Context) {
	var user model.Student
	services.OpenDatabase()
	services.Db.Where("student_number = ?", c.Param("id")).First(&user)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Schedule not found!"})
		return
	}

	var requests []RequestS
	rows, _ := services.Db.Raw("Select schedules.*, subjects.name, subjects.type,classrooms.identifier, teachers.name as Teacher from students, subscriptions, courses, subjects, classes, schedules, classrooms, teachers where students.id = ? and students.id = subscriptions.id_student and subscriptions.id_course = courses.id and subjects.id_course = courses.id and classes.id_subject = subjects.id and schedules.id_class = classes.id and schedules.id_classroom = classrooms.id and classes.id_teacher = teachers.id and date(schedules.starting_time) = current_date order by schedules.starting_time", user.Id).Rows()
	for rows.Next() {
		services.Db.ScanRows(rows, &requests)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "schedules": requests})
}

func GetStudentsByCourse(c *gin.Context) {
	var students []model.Student
	var course model.Course
	services.Db.Find(&course, c.Param("id"))

	services.OpenDatabase()
	rows, _ := services.Db.Raw("Select distinct students.* from students, subscriptions, courses where students.id = subscriptions.id_student and subscriptions.id_course = ? and subscriptions.deleted_at is null", course.Id).Rows()

	for rows.Next() {
		services.Db.ScanRows(rows, &students)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "students": students})
}
