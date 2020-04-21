package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/pilillo/go-api/users/app"
)

func main() {
	log.Info("Init server")
	app.StartApplication()
}
