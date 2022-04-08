package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func InsertSchedule(c *gin.Context) {
	controllers.InsertSchedule(c)
}
