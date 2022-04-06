package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func LoginAdmin(c *gin.Context) {
	controllers.LoginHandler(c)
}
