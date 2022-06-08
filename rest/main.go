package main

import (
	"github.com/arstrel/rest-banking/rest/app"
	"github.com/arstrel/rest-banking/rest/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
