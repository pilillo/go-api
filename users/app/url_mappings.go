package app

import (
	//log "github.com/micro/go-micro/v2/logger"

	"github.com/pilillo/go-api/users/controllers/ping"
	"github.com/pilillo/go-api/users/controllers/users"
)

func mapUrls() {
	// health check
	router.GET("/ping", ping.Ping)

	// services
	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)
}
