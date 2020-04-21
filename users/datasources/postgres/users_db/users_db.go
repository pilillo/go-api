package users_db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
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

	users_pgres_username = os.Getenv(DB_USERNAME)
	users_pgres_password = os.Getenv(DB_PASSWORD)
	users_pgres_host     = os.Getenv(DB_HOST)
	users_pgres_db       = os.Getenv(DB_SCHEMA)
)

// init is run only once when the module is imported the first time
func init() {
	dataSourceName := fmt.Sprintf(
		//"postgres://%s:%s@%s/%s?sslmode=verify-full",
		"postgres://%s:%s@%s/%s?sslmode=disable",
		users_pgres_username, users_pgres_password, users_pgres_host, users_pgres_db,
	)
	var err error
	DBCLient, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	} else {
		if err = DBCLient.Ping(); err != nil {
			panic(err)
		} else {
			log.Info("Successfully connected to db")
		}
	}

}
