package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func InsertClassroom(c *gin.Context) {
	controllers.InsertClassroom(c)
}