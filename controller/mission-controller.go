package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/repository"
)

type MissionController struct {
	Repo *repository.MissionRepository
}

func (co *MissionController) GetAllMissions(c *gin.Context) {
	missions, _ := co.Repo.GetAllMissions()
	c.JSON(200, gin.H{
		"message":  "DEBUG: GetAllMissions has been called!",
		"missions": missions,
	})
}

func (co *MissionController) GetAMission(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DEBUG: GetAMission has been called!",
		"id":      c.Param("id"),
	})
}
