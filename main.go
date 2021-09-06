package main

import (
	"hr-app-go/config"
)

func main() {
	r := config.InitRoutes()

	r.Run()
}
