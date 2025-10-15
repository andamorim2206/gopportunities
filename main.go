package main

import (
	"github.com/andamorim2206/gopportunities/config"
	"github.com/andamorim2206/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// Initializer Confings
	err := config.Init()
	if err != nil {
		logger.Errof("config initialization error: %v", err)
		return
	}

	// Initializer router
	router.Initialize()
}
