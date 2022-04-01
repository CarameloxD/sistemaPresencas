package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func GetStudentByNumber(c *gin.Context) {
	controllers.GetStudentByNumber(c)
}

func InsertStudent(c *gin.Context) {
	controllers.InsertStudent(c)
}

func InsertClassroom(c *gin.Context) {
	controllers.InsertClassroom(c)
}
