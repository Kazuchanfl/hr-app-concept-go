package config

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/controller"
)

func InitRoutes(repos Repositories) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	mc := controller.MissionController{R: repos.MissionR}
	r.GET("/missions", mc.GetAllMissions)
	r.GET("/missions/:id", mc.GetAMission)

	gc := controller.GoalController{}
	r.GET("/goals", gc.GetAllGoals)
	r.GET("/goals/:id", gc.GetAGoal)
	r.POST("/goals", gc.CreateAGoal)

	uc := controller.UserController{R: repos.UserR}
	r.GET("/users/:id", uc.GetAUser)
	r.PUT("/users/:id", uc.UpdateAUser)
	r.GET("/users/:id/goals", uc.GetUserGoals)

	return r
}
