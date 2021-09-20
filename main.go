package main

import (
	"hr-app-go/config"
)

func main() {
	db := config.InitDb()
	repos := config.InitRepositories(db)
	r := config.InitRoutes(repos)
	r.Run()
}
