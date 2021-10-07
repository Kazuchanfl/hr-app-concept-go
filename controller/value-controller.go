package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/repository"
)

type ValueController struct {
	R *repository.ValueRepository
}

func (co *ValueController) GetAllValues(c *gin.Context) {
	values, _ := co.R.GetAllValues()
	c.JSON(200, gin.H{
		"message": "DEBUG: GetAllValues has been called!",
		"values":  values,
	})
}
