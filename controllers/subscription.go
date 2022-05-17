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

func GetAllSubscriptions(c *gin.Context) {
	var subscriptions []model.Subscription
	services.OpenDatabase()
	services.Db.Select("id,id_course").Find(&subscriptions)

	if len(subscriptions) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "subscriptions not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "subscriptions": subscriptions})
}

func DeleteSubscription(c *gin.Context) {
	var registration Registration
	services.OpenDatabase()
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Error! Check All Fields"})
		return
	}

	for _, idStudent := range registration.IdStudents {
		var subscription []model.Subscription

		rows, _ := services.Db.Raw("Select * from subscriptions where id_student = ? and id_course = ?", idStudent, registration.IdCourse).Rows()
		for rows.Next() {
			services.Db.ScanRows(rows, &subscription)
			services.Db.Delete(&subscription)
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Deleted!"})
}
