package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func InsertSchedule(c *gin.Context) {
	controllers.InsertSchedule(c)
}

func GetAllSchedules(c *gin.Context) {
	controllers.GetAllSchedules(c)
}

func GetStudentsBySchedule(c *gin.Context) {
	controllers.GetStudentsBySchedule(c)
}

func GetAttendingStudentsBySchedule(c *gin.Context) {
	controllers.GetAttendingStudentsBySchedule(c)
}

func GetSchedulesByClass(c *gin.Context) {
	controllers.GetSchedulesByClass(c)
}

func DeleteSchedule(c *gin.Context) {
	controllers.DeleteSchedule(c)
}
