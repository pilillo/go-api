package app

import (
	//log "github.com/micro/go-micro/v2/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication(){
	mapUrls()
	router.Run(":8080")
}