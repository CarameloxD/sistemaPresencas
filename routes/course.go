package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func GetAllCourses(c *gin.Context) {
	controllers.GetAllCourses(c)
}

func InsertCourse(c *gin.Context) {
	controllers.InsertCourse(c)
}
