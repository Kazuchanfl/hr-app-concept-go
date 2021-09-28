package controller

import "github.com/gin-gonic/gin"

type ValueController struct {
}

func (co *ValueController) GetAllValues(c *gin.Context) {
	// TODO takahashikazuaki not yet implemetned
	c.JSON(200, gin.H{
		"message": "DEBUG: GetAllValues has been called!",
	})
}
