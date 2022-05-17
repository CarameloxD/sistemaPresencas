package routes

import (
	"github.com/gin-gonic/gin"
	"sistemaPresencas/controllers"
)

func InsertSubject(c *gin.Context) {
	controllers.InsertSubject(c)
}

func GetAllSubjects(c *gin.Context) {
	controllers.GetAllSubjects(c)
}

func DeleteSubject(c *gin.Context) {
	controllers.DeleteSubject(c)
}
