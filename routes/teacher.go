package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func LoginTeacher(c *gin.Context) {
	controllers.LoginTeacher(c)
}

func InsertTeacher(c *gin.Context) {
	controllers.InsertTeacher(c)
}

func GetTeacherInfo(c *gin.Context) {
	controllers.GetTeacherInfo(c)
}

func GetAllTeachers(c *gin.Context) {
	controllers.GetAllTeachers(c)
}

func DeleteTeacher(c *gin.Context) {
	controllers.DeleteTeacher(c)
}

func GetSchedulesByTeacher(c *gin.Context) {
	controllers.GetSchedulesByTeacher(c)
}
