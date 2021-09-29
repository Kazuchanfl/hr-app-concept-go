package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/model"
	"hr-app-go/repository"
)

type UserController struct {
	R *repository.UserRepository
}

func (co *UserController) GetAUser(c *gin.Context) {
	user, err := co.R.GetAUser(c.Param("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "DEBUG: GetAUser has been called, but something went wrong.",
			"user":    nil,
			"error":   err,
		})
	}
	c.JSON(200, gin.H{
		"message": "DEBUG: GetAUser has been called!",
		"user":    user,
	})
}

func (co *UserController) UpdateAUser(c *gin.Context) {
	// TODO takahashikazuaki handle error for failing GetAUser
	user, _ := co.R.GetAUser(c.Param("id"))
	user.Name = c.PostForm("name")
	user.Profile = c.PostForm("profile")
	updatedUser, _ := co.R.UpdateAUser(user)
	c.JSON(200, gin.H{
		"message": "DEBUG: UpdateAUser has been called!",
		"user":    updatedUser,
	})
}

func (co *UserController) GetUserGoals(c *gin.Context) {
	// TODO takahashikazuaki not yet implemented
	var user model.User
	var goals []model.Goal
	c.JSON(200, gin.H{
		"message": "DEBUG: UpdateAUser has been called!",
		"user":    user,
		"goals":   goals,
	})
}
