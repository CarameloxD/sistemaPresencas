package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func InsertSubscription(c *gin.Context) {
	controllers.InsertSubscription(c)
}