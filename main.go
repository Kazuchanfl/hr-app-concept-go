package main

import (
	"hr-app-go/config"
)

func main() {
	config.InitDb()

	r := config.InitRoutes()
	r.Run()
}
