package users

import (
	"database/sql"
	"fmt"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/pilillo/go-api/users/datasources/postgres/users_db"
	"github.com/pilillo/go-api/users/utils/date"
	"github.com/pilillo/go-api/users/utils/errors"
)

const (

	// queries
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?);"
	querySelectUser             = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

func (user *User) Get() *errors.RestErr {
	var stmt *sql.Stmt = nil
	var resErr *errors.RestErr = nil

	stmt, err := users_db.DBCLient.Prepare(querySelectUser)
	if err != nil {
		logger.Error("Error while preparing user statement :: ", err)
		resErr = errors.GetInternalServerError(err.Error())
	} else {
		defer stmt.Close()
		row := stmt.QueryRow(user.Id) // WHERE clause
		if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.Status); err != nil {
			log.Error(err.Error())
			resErr = errors.GetInternalServerError(fmt.Sprintf("Error while accessing user %d", user.Id))
		}
	}
	return resErr
}

func (user *User) Save() *errors.RestErr {
	var result *errors.RestErr

	stmt, err := users_db.DBCLient.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error while preparing user statement :: ", err)
		result = errors.GetInternalServerError(err.Error())
	} else {
		defer stmt.Close() // close connection upon function ends

		user.DateCreated = date.GetNowString()

		// execute statement
		insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

		if err != nil {
			result = errors.GetInternalServerError(fmt.Sprintf("Error when trying to save user :: %s", err.Error()))
		} else {
			userId, err := insertResult.LastInsertId()
			if err != nil {
				result = errors.GetInternalServerError(fmt.Sprintf("Error when trying to save user :: %s", err.Error()))
			}
			user.Id = userId
		}
	}
	return result
}

func (user *User) Update() *errors.RestErr {
	var result *errors.RestErr
	stmt, err := users_db.DBCLient.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error while preparing user statement :: ", err)
		result = errors.GetInternalServerError(err.Error())
	} else {
		defer stmt.Close() // close connection upon function ends
		_, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
		if err != nil {
			// todo: parse error
			result = errors.GetInternalServerError(err.Error())
		}
	}
	return result
}

func (user *User) Delete() *errors.RestErr {
	var result *errors.RestErr
	stmt, err := users_db.DBCLient.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error while preparing user statement :: ", err)
		result = errors.GetInternalServerError(err.Error())
	} else {
		defer stmt.Close()
		_, err := stmt.Exec(user.Id)
		if err != nil {
			// todo: parse+handle error properly
			result = errors.GetInternalServerError(err.Error())
		}
	}
	return result
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	var resErr *errors.RestErr
	var resUsers []User

	stmt, err := users_db.DBCLient.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("Error while preparing user statement :: ", err)
		resErr = errors.GetInternalServerError(err.Error())
	} else {
		defer stmt.Close()
		rows, err := stmt.Query(status)

		if err != nil {
			resErr = errors.GetInternalServerError(err.Error())
		} else {
			defer rows.Close()
			resUsers = make([]User, 0)
			for rows.Next() {
				var user User
				// populate newly created user with result from select
				if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
					resErr = errors.GetInternalServerError(err.Error())
				} else {
					resUsers = append(resUsers, user)
				}
			}

			if len(resUsers) == 0 {
				resErr = errors.GetNotFoundError(fmt.Sprintf("no users matching the status %s", status))
				resUsers = nil
			}
		}
	}

	return resUsers, resErr
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	var stmt *sql.Stmt = nil
	var resErr *errors.RestErr = nil

	stmt, err := users_db.DBCLient.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("Error while preparing user statement :: ", err)
		resErr = errors.GetInternalServerError(err.Error())
	} else {
		defer stmt.Close()
		row := stmt.QueryRow(user.Email, user.Password, StatusActive)
		if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.Status); err != nil {
			log.Error(fmt.Sprintf("error while trying to get user :: %s", err.Error()))
			resErr = errors.GetInternalServerError(fmt.Sprintf("Error while accessing user %d", user.Id))
		}
	}
	return resErr
}
