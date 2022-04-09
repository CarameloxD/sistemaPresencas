package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func InsertAttendance(c *gin.Context) {
	controllers.InsertAttendance(c)
}
