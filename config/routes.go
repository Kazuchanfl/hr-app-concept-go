package config

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/controller"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	missionController := controller.MissionController{}

	r.GET("/missions", missionController.GetAllMissions)

	return r
}
