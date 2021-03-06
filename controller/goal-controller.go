package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/model"
)

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

func (controller *GoalController) GetAGoal(c *gin.Context) {
	// TODO get a Goal from DB
	goal := model.Goal{
		Title:       "頑張る",
		Description: "めっちゃ頑張る",
	}
	c.JSON(200, gin.H{
		"message": "DEBUG: GetAGoal has been called!",
		"id":      c.Param("id"),
		"goal":    goal,
	})
}

func (controller *GoalController) CreateAGoal(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	// TODO takahashikazuaki Store goal data in DB
	goal := model.Goal{
		Title:       title,
		Description: description,
	}

	c.JSON(200, gin.H{
		"message": "DEBUG: CreateAGoal has been called!",
		"goal":    goal,
	})
}
