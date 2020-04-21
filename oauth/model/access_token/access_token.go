package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/pilillo/go-api/oauth/utils/errors"
	"github.com/pilillo/go-api/users/utils/crypto"
)

const (
	expirationTime = 24
	// supported grant types
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// used for password grant_type
	Email    string `json:"email"`
	Password string `json:"password"`

	// used for client credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	var resErr *errors.RestErr

	switch at.GrantType {
	case grantTypePassword:
		break

	case grantTypeClientCredentials:
		break

	default:
		resErr = errors.GetBadRequestError("invalid grant_type parameter")
	}

	return resErr
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	var resErr *errors.RestErr
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		resErr = errors.GetBadRequestError("invalid access token id")
	}

	if at.UserId <= 0 {
		resErr = errors.GetBadRequestError("invalid user id")
	}

	if at.ClientId <= 0 {
		resErr = errors.GetBadRequestError("invalid client id")
	}

	if at.Expires <= 0 {
		resErr = errors.GetBadRequestError("invalid expiration time")
	}
	return resErr
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}
func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return expirationTime.Before(now)
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto.GetSha512(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
