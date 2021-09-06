package controller

import "github.com/gin-gonic/gin"

type MissionRequestHandler interface {
	GetAllMissions(*gin.Context)
}

type MissionController struct {
}

func (controller *MissionController) GetAllMissions(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAllMission has been called!",
	})
}
