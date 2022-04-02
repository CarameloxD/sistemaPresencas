package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func InsertClass(c *gin.Context) {
	controllers.InsertClass(c)
}