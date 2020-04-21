package users

import (
	//log "github.com/micro/go-micro/v2/logger"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pilillo/go-api/users/model/users"
	"github.com/pilillo/go-api/users/services"
	"github.com/pilillo/go-api/users/utils/errors"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.GetBadRequestError("invalid user id, it should be a number")
	} else {
		return userId, nil
	}
}

func Create(c *gin.Context) {
	var user users.User
	// unmarshal the incoming string to a json
	// attempt to bind the json to user struct definition
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.GetBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status, restErr)
	} else {
		// call service to add the user
		result, saveErr := services.UsersService.CreateUser(user)
		if saveErr != nil {
			c.JSON(saveErr.Status, saveErr)
		} else {
			c.JSON(http.StatusCreated, result)
		}
	}

}

func Get(c *gin.Context) {
	//var err *errors.RestErr
	// parse user id as int64 base 10
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	} else {
		user, getErr := services.UsersService.GetUser(userId)
		if getErr != nil {
			c.JSON(getErr.Status, getErr)
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	} else {
		var user users.User
		// validate incoming json with the user struct
		if err := c.ShouldBindJSON(&user); err != nil {
			restErr := errors.GetBadRequestError("invalid json body")
			c.JSON(restErr.Status, restErr)
		} else {
			// set user with passed user
			user.Id = userId

			// attempt updating user
			res, err := services.UsersService.UpdateUser(user)
			if err != nil {
				c.JSON(err.Status, res)
			} else {
				c.JSON(http.StatusOK, res)
			}
		}
	}
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	} else {
		if err := services.UsersService.DeleteUser(userId); err != nil {
			c.JSON(err.Status, err)
		} else {
			c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
		}
	}
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.GetBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
	} else {
		user, err := services.UsersService.LoginUser(request)
		if err != nil {
			c.JSON(err.Status, err)
		} else {
			c.JSON(http.StatusOK, user)
			//c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public") == "true"))
		}
	}
}
