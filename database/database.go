package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:kraken1288@tcp(localhost:3306)/golang_session")
	if err != nil {
		panic(err)
	}
	return db
}
