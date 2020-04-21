package services

import (
	"github.com/pilillo/go-api/users/model/users"
	"github.com/pilillo/go-api/users/utils/crypto"
	"github.com/pilillo/go-api/users/utils/date"
	"github.com/pilillo/go-api/users/utils/errors"
)

// static global var UserService is used to group all methods of type userServiceInterface
var (
	UsersService usersServiceInterface = &usersService{}
)

// group all methods under the userService type
// i.e. all methods are applied to the type userService (aka class in OOP)
type usersService struct{}

// the userService implements all methods defined in the userServiceInterface
// this way we can apply polymorphism
type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) ([]users.User, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	var resUsr *users.User
	var resErr *errors.RestErr
	if err := user.Validate(); err != nil {
		resErr = err
	} else {
		user.Status = users.StatusActive
		user.DateCreated = date.GetNowString()
		user.Password = crypto.GetSha512(user.Password)
		if err := user.Save(); err != nil {
			resErr = err
		} else {
			resUsr = &user
		}

	}
	return resUsr, resErr

}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	var err *errors.RestErr
	if err := result.Get(); err != nil {
		result = nil
	}
	return result, err

}

func (s *usersService) UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	var resVal *users.User
	var resErr *errors.RestErr
	currUser, err := UsersService.GetUser(user.Id)
	if err != nil {
		resErr = err
	} else {
		// validate the user
		if err := user.Validate(); err != nil {
			resErr = err
		} else {
			currUser.FirstName = user.FirstName
			currUser.LastName = user.LastName
			currUser.Email = user.Email
			if err := currUser.Update(); err != nil {
				resErr = err
			} else {
				resVal = currUser
			}
		}

	}
	return resVal, resErr
}

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	dao := &users.User{Id: userId}
	return dao.Delete()
}

func (s *usersService) SearchUser(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto.GetSha512(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
