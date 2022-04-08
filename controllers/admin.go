package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"

	"net/http"
	"sistemaPresencas/model"
	"sistemaPresencas/services"
)

func LoginHandler(c *gin.Context) {
	var creds model.Admin
	var usr model.Admin

	if err := c.ShouldBindJSON(&creds); err != nil { // guardo no creds o que veio por parametro no pedido
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	fmt.Println(creds)
	fmt.Println(creds.Id)
	services.OpenDatabase()
	services.Db.Find(&usr, "username = ?", creds.Username) //procuro na bd

	if usr.Username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		return
	} else { //
		token := services.GenerateTokenJWT(usr) // gero o token

		if token == "" { // se o token ta vazio, acesso negado
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Access denied!"})
			return
		}
		//defer services.Db.Close()

		// caso td de certo..
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "token": token, "ID": usr.ID, "username": usr.Name})
		fmt.Println(usr.Name)
	}
}
