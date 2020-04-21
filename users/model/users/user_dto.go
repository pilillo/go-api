package users

import (
	"strings"

	"github.com/pilillo/go-api/users/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-"` // do not serialize/deserialize to json (-)
}

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	if user.FirstName == "" {
		return errors.GetBadRequestError("invalid first name")
	}

	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	if user.LastName == "" {
		return errors.GetBadRequestError("invalid last name")
	}

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.GetBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(strings.ToLower(user.Password))
	if user.Password == "" {
		return errors.GetBadRequestError("invalid password")
	}

	return nil
}
