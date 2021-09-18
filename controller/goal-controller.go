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
	// TODO takahashikazuaki Followings are example codes so they could be deleted.
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})

	//c.JSON(200, gin.H{
	//	"message": "DEBUG: CreateAGoal has been called!",
	//	//"goal":    goal,
	//})
}
