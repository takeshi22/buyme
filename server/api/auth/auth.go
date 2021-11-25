package controllers

import (
	"log"

	"github.com/buyme/models"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)

		c.JSON(400, gin.H{"msg": "could not bind user"})
		c.Abort()
		return
	}

	err = user.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())

		c.JSON(500, gin.H{"msg": "error encrypt password"})
		c.Abort()
		return
	}

	err = user.CreateUser()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"msg": "could not create user"})
		c.Abort()
		return
	}

	c.JSON(200, user)
}

func Login(c *gin.Context) {

}
