package go_db_mysequel

import (
	"testing"
	
	"database/sql"
	// without this driver unknown
	_ "github.com/go-sql-driver/mysql" // init
)

func TestOpenConnection(t *testing.T)  {
	// DB Connection
	db, err := sql.Open("mysql", "root:mysqlrootpass@tcp(localhost:3306)/belajar_go_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Use DB

}