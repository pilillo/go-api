package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pilillo/go-api/oauth/http"
	"github.com/pilillo/go-api/oauth/repository/db"
	"github.com/pilillo/go-api/oauth/repository/rest"
	"github.com/pilillo/go-api/oauth/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	/*
		session, dbErr := cassandra.GetSession()
		if dbErr != nil {
			panic(dbErr)
		}
	*/
	//session := cassandra.GetSession()
	//session.Close()

	//atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewAccessTokenHandler(
		access_token.NewService(
			rest.NewRestUsersRepository(),
			db.NewRepository(),
		),
	)

	// set routes
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token/", atHandler.Create)

	router.Run(":8080")
}
