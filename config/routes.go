package config

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/controller"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	missionController := controller.MissionController{}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/missions", missionController.GetAllMissions)
	r.GET("/missions/:id", missionController.GetAMission)

	return r
}
