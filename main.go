package main

import (
	"github.com/arstrel/rest-banking/app"
	"github.com/arstrel/rest-banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
