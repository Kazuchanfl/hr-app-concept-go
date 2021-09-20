package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/repository"
)

type MissionController struct {
	Repository *repository.MissionRepository
}

func (controller *MissionController) GetAllMissions(c *gin.Context) {
	missions, _ := controller.Repository.GetAllMissions()
	c.JSON(200, gin.H{
		"message":  "DEBUG: GetAllMissions has been called!",
		"missions": missions,
	})
}

func (controller *MissionController) GetAMission(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DEBUG: GetAMission has been called!",
		"id":      c.Param("id"),
	})
}
