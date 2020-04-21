package db

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/pilillo/go-api/oauth/clients/cassandra"
	"github.com/pilillo/go-api/oauth/model/access_token"
	"github.com/pilillo/go-api/oauth/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?,?,?,?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct{}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (service *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	var resVal *access_token.AccessToken
	var resErr *errors.RestErr

	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(
		queryGetAccessToken,
		id,
	).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		//if err == gocql.ErrNotFound {
		// same as below, although as string comparison
		if err.Error() == gocql.ErrNotFound.Error() {
			resErr = errors.GetNotFoundError(fmt.Sprintf("no access token found with id %s", id))
		} else {
			resErr = errors.GetInternalServerError(err.Error())
		}
	} else {
		resVal = &result
	}

	return resVal, resErr
}

func (service *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	var resErr *errors.RestErr

	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		resErr = errors.GetInternalServerError(err.Error())
	}

	return resErr
}

func (service *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	var resErr *errors.RestErr

	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		resErr = errors.GetInternalServerError(err.Error())
	}

	return resErr
}
