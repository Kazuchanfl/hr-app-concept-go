package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/model"
	"hr-app-go/repository"
)

type MissionController struct {
	Repository *repository.MissionRepository
}

func (controller *MissionController) GetAllMissions(c *gin.Context) {
	// TODO get Missions from DB
	missions := []model.Mission{
		{Statement: "人々を技術で前進させる"},
		{Statement: "生き物への慈愛の念を広める"},
	}
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
