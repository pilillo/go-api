package access_token

import (
	"strings"

	"github.com/pilillo/go-api/oauth/model/access_token"
	"github.com/pilillo/go-api/oauth/repository/db"
	"github.com/pilillo/go-api/oauth/repository/rest"
	"github.com/pilillo/go-api/oauth/utils/errors"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
}

func NewService(usersRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	var resVal *access_token.AccessToken
	var resErr *errors.RestErr

	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		resErr = errors.GetBadRequestError("invalid access token id")
	} else {
		accessToken, err := s.dbRepo.GetById(accessTokenId)
		if err != nil {
			resErr = err
		} else {
			resVal = accessToken
		}
	}
	return resVal, resErr
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	var resVal *access_token.AccessToken
	var resErr *errors.RestErr

	if err := request.Validate(); err != nil {
		resErr = err
	} else {
		// authenticate the user against the users API
		user, err := s.restUsersRepo.LoginUser(request.Email, request.Password)
		if err != nil {
			resErr = err
		} else {
			// generate a new access token
			at := access_token.GetNewAccessToken(user.Id)
			at.Generate() // todo: WTF is this?

			if err := s.dbRepo.Create(at); err != nil {
				resErr = err
			} else {
				resVal = &at
			}
		}
	}

	return resVal, resErr
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	var resErr *errors.RestErr
	if err := at.Validate(); err != nil {
		resErr = err
	} else {
		resErr = s.dbRepo.UpdateExpirationTime(at)
	}
	return resErr
}
