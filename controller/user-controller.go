package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/repository"
)

type UserController struct {
	UR *repository.UserRepository
	GR *repository.GoalRepository
}

func (co *UserController) GetAUser(c *gin.Context) {
	user, err := co.UR.GetAUser(c.Param("id"))
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
	user, _ := co.UR.GetAUser(c.Param("id"))
	user.Name = c.PostForm("name")
	user.Profile = c.PostForm("profile")
	updatedUser, _ := co.UR.UpdateAUser(user)
	c.JSON(200, gin.H{
		"message": "DEBUG: UpdateAUser has been called!",
		"user":    updatedUser,
	})
}

func (co *UserController) GetUserGoals(c *gin.Context) {
	user, _ := co.UR.GetAUser(c.Param("id"))
	goals, _ := co.GR.GetGoalsOfAUser(string(user.ID))
	c.JSON(200, gin.H{
		"message": "DEBUG: GetUserGoals has been called!",
		"user":    user,
		"goals":   goals,
	})
}
