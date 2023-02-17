package controller

import (
	"parkezy/server/model"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	// get data from request body
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// create user
	user := model.User{Email: input.Email, Password: input.Password}
	model.DB.Create(&user)
	c.JSON(200, gin.H{"data": user})
}

func LoginUser(c *gin.Context) {
	// get data from request body
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// get user
	var user model.User
	if err := model.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}
	// validate password
	if input.Password != user.Password {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	c.JSON(200, gin.H{"data": user})
}
