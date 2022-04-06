package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
)

type Registration struct {
	IdCourse   int
	IdStudents []int
}

func InsertSubscription(c *gin.Context) {
	var registration Registration
	services.OpenDatabase()
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}

	for _, idStudent := range registration.IdStudents {
		var subscription model.Subscription
		subscription.IdCourse = registration.IdCourse
		subscription.IdStudent = idStudent
		services.Db.Save(&subscription)
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Created Successfully"})
}
