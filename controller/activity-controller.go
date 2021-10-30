package controller

import (
	"github.com/gin-gonic/gin"
	"hr-app-go/model"
	"hr-app-go/repository"
)

type ActivityController struct {
	AR *repository.ActivityRepository
	UR *repository.UserRepository
}

func (co *ActivityController) CreateAnActivity(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	user, _ := co.UR.GetAUser(c.Param("id"))

	activity := model.Activity{
		Title:       title,
		Description: description,
		UserID:      user.ID,
	}
	_ = co.AR.CreateAnActivity(activity)

	c.JSON(200, gin.H{
		"message":  "DEBUG: CreateAnActivity has been called!",
		"activity": activity,
	})
}
