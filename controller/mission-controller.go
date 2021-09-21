package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/repository"
)

type MissionController struct {
	R *repository.MissionRepository
}

func (co *MissionController) GetAllMissions(c *gin.Context) {
	missions, _ := co.R.GetAllMissions()
	c.JSON(200, gin.H{
		"message":  "DEBUG: GetAllMissions has been called!",
		"missions": missions,
	})
}

func (co *MissionController) GetAMission(c *gin.Context) {
	mission := co.R.GetAMission(c.Param("id"))
	c.JSON(200, gin.H{
		"message": "DEBUG: GetAMission has been called!",
		"mission": mission,
	})
}
