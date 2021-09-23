package controller

import (
	"github.com/gin-gonic/gin"
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
