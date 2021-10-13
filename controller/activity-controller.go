package controller

import "github.com/gin-gonic/gin"

// TODO takahashikazuaki ActivityController should be included
type ActivityController struct {
}

func (co *ActivityController) CreateAnActivity(c *gin.Context) {
	// TODO takahashikazuaki store an activity in DB
	c.JSON(200, gin.H{
		"message": "DEBUG: CreateAnActivity has been called!",
	})
}
