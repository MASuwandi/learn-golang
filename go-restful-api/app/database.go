package app

import (
	"database/sql"
	"go-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:mysqlrootpass@tcp(localhost:3306)/go_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}