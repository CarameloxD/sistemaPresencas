package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func InsertSubscription(c *gin.Context) {
	controllers.InsertSubscription(c)
}

func DeleteSubscription(c *gin.Context) {
	controllers.DeleteSubscription(c)
}

func GetAllSubscriptions(c *gin.Context) {
	controllers.GetAllSubscriptions(c)
}
