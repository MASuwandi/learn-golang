package go_db_mysequel

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	// parse time, time convert
	parseTime := "?parseTime=true"
	db, err := sql.Open("mysql", "root:mysqlrootpass@tcp(localhost:3306)/belajar_go_db"+parseTime)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
