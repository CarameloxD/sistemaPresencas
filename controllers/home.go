package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
)

func GetStudentByNumber(c *gin.Context) {
	var user model.Student
	services.OpenDatabase()
	services.Db.Find(&user, c.Param("student_number"))

	if user.StudentNumber == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}
