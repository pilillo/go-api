package rest

import (
	"encoding/json"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/pilillo/go-api/oauth/model/users"
	"github.com/pilillo/go-api/oauth/utils/errors"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://something.com", // todo: change domain
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {
}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	var resVal *users.User
	var resErr *errors.RestErr

	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	// call the login REST service exposed within the network (private)
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.GetInternalServerError("invalid restclient response while trying to login user")
	} else if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			resErr = errors.GetInternalServerError("invalid error interface when trying to login user")
		}
	} else {
		var user users.User
		if err := json.Unmarshal(response.Bytes(), &user); err != nil {
			resErr = errors.GetInternalServerError("error when trying to unmarshal users response from http service")
		} else {
			resVal = &user
		}
	}
	return resVal, resErr
}
