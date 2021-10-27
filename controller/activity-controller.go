package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/repository"
)

type ActivityController struct {
	R *repository.ActivityRepository
}

func (co *ActivityController) CreateAnActivity(c *gin.Context) {
	// TODO takahashikazuaki create an activity instance and pass it to repository
	_ = co.R.CreateAnActivity()
	c.JSON(200, gin.H{
		"message": "DEBUG: CreateAnActivity has been called!",
	})
}
