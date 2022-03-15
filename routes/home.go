package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func GetStudentByNumber(c *gin.Context) {
	controllers.GetStudentByNumber(c)
}
