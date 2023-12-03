package main

import (
	"github.com/leobelini-studies/arthur-404-dev_api_moderna_golang/config"
	"github.com/leobelini-studies/arthur-404-dev_api_moderna_golang/router"
)

var (
	logger *config.Logger
)

func main() {

	// Initialize Logger
	logger = config.NewLogger("main")

	// Initialize Config
	err := config.Init()

	if err != nil {
		logger.Errorf("config.Init() error: %s", err)
		panic(err)
	}

	// Initialize Router
	router.Initialize()
}

// Qualquer função que inicia com letra maiúscula é exportada.
