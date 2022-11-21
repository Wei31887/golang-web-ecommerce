package data

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	err error
	Db *sql.DB
)

func init() {
	Db, err = sql.Open("postgres", "dbname=store sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
}

