package main

import (
	"github.com/sjdpk/bankingapp/app"
	"github.com/sjdpk/bankingapp/logger"
)

func main() {
	logger.Info("Server is starting ")
	app.Start()
}
