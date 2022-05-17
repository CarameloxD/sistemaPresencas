package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func InsertClass(c *gin.Context) {
	controllers.InsertClass(c)
}

func GetAllClasses(c *gin.Context) {
	controllers.GetAllClasses(c)
}

func DeleteClass(c *gin.Context) {
	controllers.DeleteClass(c)
}
