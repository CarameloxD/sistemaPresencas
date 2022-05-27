package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func GetQrCode(c *gin.Context) {
	controllers.GetQrCode(c)
}
