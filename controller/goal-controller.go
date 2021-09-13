package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/model"
)

type GoalRequestHandler interface {
	GetAllGoals(*gin.Context)
	GetAGoal(*gin.Context)
}

type GoalController struct {
}

func (controller *GoalController) GetAllGoals(c *gin.Context) {
	// TODO get Goals from DB
	goals := []model.Goal{
		{Title: "頑張る", Description: "めっちゃ頑張る"},
		{Title: "もっと頑張る", Description: "すごい頑張る"},
	}
	c.JSON(200, gin.H{
		"message": "DEBUG: GetAllGoals has been called!",
		"goals":   goals,
	})
}
