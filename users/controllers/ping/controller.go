package ping

import (
	"github.com/gin-gonic/gin"
	//log "github.com/micro/go-micro/v2/logger"
	"net/http"
)

func Ping(c *gin.Context){
	c.String(http.StatusOK, "pong")
}