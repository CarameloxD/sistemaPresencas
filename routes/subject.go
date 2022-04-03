package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func InsertSubject(c *gin.Context) {
	controllers.InsertSubject(c)
}
