package users_db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/common/log"
)

const (
	DB_USERNAME = "DB_USERNAME"
	DB_PASSWORD = "DB_PASSWORD"
	DB_HOST     = "DB_HOST"
	DB_SCHEMA   = "DB_SCHEMA"
)

var (
	DBCLient *sql.DB

	users_mysql_username = os.Getenv(DB_USERNAME)
	users_mysql_password = os.Getenv(DB_PASSWORD)
	users_mysql_host     = os.Getenv(DB_HOST)
	users_mysql_schema   = os.Getenv(DB_SCHEMA)
)

// init is run only once when the module is imported the first time
func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		users_mysql_username, users_mysql_password, users_mysql_host, users_mysql_schema,
	)
	var err error
	DBCLient, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	} else {
		if err = DBCLient.Ping(); err != nil {
			panic(err)
		} else {
			log.Info("Successfully connected to db")
		}
	}

}
