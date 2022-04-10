package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func GetStudentByNumber(c *gin.Context) {
	controllers.GetStudentByNumber(c)
}

func InsertStudent(c *gin.Context) {
	controllers.InsertStudent(c)
}

func GetAllStudents(c *gin.Context) {
	controllers.GetAllStudents(c)
}

func DeleteStudent(c *gin.Context) {
	controllers.DeleteStudent(c)
}

func GetSchedulesByStudent(c *gin.Context) {
	controllers.GetSchedulesByStudent(c)
}
