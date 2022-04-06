package controllers

import (
	"fmt"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
	"github.com/gin-gonic/gin"
)

type Registration struct {
	idCourse   int
	idStudents []int
}

func InsertSubscription(c *gin.Context) {
	services.OpenDatabase()
	var registration Registration
	/*if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}
	*/
	fmt.Print(registration.idCourse)
	var courseId = registration.idCourse
	for _, student := range registration.idStudents {
		var subscription model.Subscription
		subscription.IdCourse = courseId
		subscription.IdStudent = student
		services.Db.Save(&subscription)
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})
}
