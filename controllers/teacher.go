package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
	//"time"
)

func LoginTeacher(c *gin.Context) {
	var usr model.Teacher
	fmt.Println(c.Params)
	if err := c.ShouldBindJSON(&usr); err != nil { // guardo no creds o que veio por parametro no pedido
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	fmt.Println(usr)
	fmt.Println(usr.Id)
	services.OpenDatabase()
	services.Db.Find(&usr, "username = ?", usr.Username) //procuro na bd

	if usr.Username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "ID": usr.ID, "username": usr.Name})
		fmt.Println(usr.Name)
	}
}

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

func GetAllTeachers(c *gin.Context) {
	var teachers []model.Teacher

	services.OpenDatabase()
	services.Db.Select("Id,name,username,email").Find(&teachers)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "teachers": teachers})
}

func DeleteTeacher(c *gin.Context) {
	var user model.Teacher
	services.OpenDatabase()
	services.Db.Find(&user, c.Param("Id"))

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}
	services.Db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Deleted!"})
}

func GetSchedulesByTeacher(c *gin.Context) {
	var user model.Teacher
	services.OpenDatabase()
	services.Db.Where("username = ?", c.Param("id")).First(&user)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Schedule not found!"})
		return
	}

	var requests []RequestS
	rows, _ := services.Db.Raw("Select schedules.*, subjects.name, subjects.type,classrooms.identifier, teachers.name as Teacher from subjects, classes, schedules, classrooms, teachers where teachers.username = ? and classes.id_subject = subjects.id and schedules.id_class = classes.id and schedules.id_classroom = classrooms.id and classes.id_teacher = teachers.id and date(schedules.starting_time) = current_date order by schedules.starting_time", user.Username).Rows()
	for rows.Next() {
		services.Db.ScanRows(rows, &requests)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "schedules": requests})
}
