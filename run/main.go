package main

import (
	"bizCard/router"
)

func main() {
	engine := router.SetupRouter()
	err := engine.Run()
	if err != nil {
		return
	}
}
