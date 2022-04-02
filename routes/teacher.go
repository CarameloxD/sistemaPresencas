package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func InsertTeacher(c *gin.Context){
	controllers.InsertTeacher(c)
}

func GetTeacherInfo(c *gin.Context){
	controllers.GetTeacherInfo(c)
}