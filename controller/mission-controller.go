package controller

import "github.com/gin-gonic/gin"

type MissionRequestHandler interface {
	GetAllMissions(*gin.Context)
	GetAMission(*gin.Context)
}

type MissionController struct {
}

func (controller *MissionController) GetAllMissions(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAllMissions has been called!",
	})
}

func (controller *MissionController) GetAMission(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAMission has been called!",
		"id":      c.Param("id"),
	})
}
