package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func GetUserById(c *gin.Context) {
	controllers.GetUserById(c)
}
