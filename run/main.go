package main

import (
	"bizCard/router"
)

func main() {
	router.SetupService()
	engine := router.SetupRouter()
	err := engine.Run()
	if err != nil {
		return
	}
}
