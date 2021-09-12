package controller

import "github.com/gin-gonic/gin"

type GoalRequestHandler interface {
	GetAllGoals(*gin.Context)
	GetAGoal(*gin.Context)
}

type GoalController struct {
}
