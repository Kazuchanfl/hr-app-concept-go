package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/model"
)

type UserController struct {
	// TODO takahashikazuaki use UserRepository in UserController
	//R *repository.UserRepository
}

func (co *UserController) GetAUser(c *gin.Context) {
	// TODO takahashikazuaki get a user from DB
	user := model.User{
		Name:    "Kazuaki Takahashi",
		Profile: "Software Developer at teamLab Inc.",
	}
	c.JSON(200, gin.H{
		"message": "DEBUG: GetAUser has been called!",
		"user":    user,
	})
}
