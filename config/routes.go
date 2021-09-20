package config

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/controller"
	"hr-app-go/repository"
)

func InitRoutes(missionRepository *repository.MissionRepository) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	missionController := controller.MissionController{Repository: missionRepository}
	r.GET("/missions", missionController.GetAllMissions)
	r.GET("/missions/:id", missionController.GetAMission)

	goalController := controller.GoalController{}
	r.GET("/goals", goalController.GetAllGoals)
	r.GET("/goals/:id", goalController.GetAGoal)
	r.POST("/goals", goalController.CreateAGoal)

	return r
}
