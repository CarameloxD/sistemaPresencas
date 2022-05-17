package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func InsertClassroom(c *gin.Context) {
	controllers.InsertClassroom(c)
}

func GetAllClassrooms(c *gin.Context) {
	controllers.GetAllClassrooms(c)
}

func DeleteClassroom(c *gin.Context) {
	controllers.DeleteClassroom(c)
}
