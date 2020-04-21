package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	atDomain "github.com/pilillo/go-api/oauth/model/access_token"
	"github.com/pilillo/go-api/oauth/services/access_token"
	"github.com/pilillo/go-api/oauth/utils/errors"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Servicels
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
	} else {
		c.JSON(http.StatusOK, accessToken)
	}
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.GetBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
	} else {
		if at, err := handler.service.Create(at); err != nil {
			c.JSON(err.Status, err)
		} else {
			c.JSON(http.StatusCreated, at)
		}
	}
}
