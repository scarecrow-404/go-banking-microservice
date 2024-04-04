package main

import (
	"github.com/scarecrow-404/banking/app"
	"github.com/scarecrow-404/banking/logger"
)

func main() {
	logger.Info(". . . . Starting the application . . . .")
	app.Start()
}