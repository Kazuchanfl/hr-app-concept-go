package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/repository"
)

type ValueController struct {
	R *repository.ValueRepository
}

func (co *ValueController) GetAllValues(c *gin.Context) {
	// TODO takahashikazuaki not yet implemetned
	c.JSON(200, gin.H{
		"message": "DEBUG: GetAllValues has been called!",
	})
}
