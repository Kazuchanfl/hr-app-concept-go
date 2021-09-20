package main

import (
	"hr-app-go/config"
	"hr-app-go/repository"
)

func main() {
	db := config.InitDb()

	missionRepository := &repository.MissionRepository{DB: db}

	r := config.InitRoutes(missionRepository)
	r.Run()
}
