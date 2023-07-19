package go_db_mysequel

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Exec SQL
func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES('thor', 'Thor')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

// Query SQL
func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `SELECT id, name
			  FROM customer`
	rows, err := db.QueryContext(ctx, query)
	// return pointer
	if err != nil {
		panic(err)
	}
	// Rows
	fmt.Println("rows: ", rows)
	for rows.Next() {
		var id, name string
		// order base on query
		// use pointer to catch result
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
	}

	defer rows.Close()

	fmt.Println("Success query customer")
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `SELECT id, name, email, balance, rating, birth_date, married, created_at
			  FROM customer`
	rows, err := db.QueryContext(ctx, query)
	// return pointer
	if err != nil {
		panic(err)
	}
	// Rows
	fmt.Println("rows: ", rows)
	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthdate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthdate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("--------------------------")
		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
		if email.Valid {
			fmt.Println("email: ", email)
		}
		fmt.Println("balance: ", balance)
		fmt.Println("rating: ", rating)
		if birthdate.Valid {
			fmt.Println("birthdate: ", birthdate)
		}
		fmt.Println("married: ", married)
		fmt.Println("createdAt: ", createdAt)
	}

	defer rows.Close()

	fmt.Println("Success query customer")
}

// SQL Injection
// bad practice
func TestSqlInjec(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Scenario 1
	// username := "admin"
	// password := "admin"

	// Scenario 2 - SQL Injection
	// don't trust user input
	// "#" equal to comment
	// string concat bad practice
	username := "admin'; #"
	password := "salah"

	query := `SELECT username
			  FROM user
			  WHERE username = '` + username +
		`' AND password = '` + password + `' LIMIT 1`

	fmt.Println("query: ", query)

	rows, err := db.QueryContext(ctx, query)
	// return pointer
	if err != nil {
		panic(err)
	}
	// Rows
	fmt.Println("rows: ", rows)
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("username: ", username)
	} else {
		fmt.Println("gagal login")
	}

	defer rows.Close()

	fmt.Println("Success query customer")
}

// SQL with Parameters
// best practice
func TestSqlInjecSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Scenario 1
	username := "admin"
	password := "admin"

	// // Scenario 2 - SQL Injection
	// username := "admin'; #"
	// password := "salah"

	query := `SELECT username
			  FROM user
			  WHERE username = ?
			  AND password = ?
			  LIMIT 1`

	rows, err := db.QueryContext(ctx, query, username, password)
	// return pointer
	if err != nil {
		panic(err)
	}
	// Rows
	fmt.Println("rows: ", rows)
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("username: ", username)
	} else {
		fmt.Println("gagal login")
	}

	defer rows.Close()

	fmt.Println("Success query customer")
}

func TestExecSqlParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "thor'; DROP TABLE user; #"
	password := "thor"

	query := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

// Auto Increment
func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "boy@mail.com"
	comment := "boyya"

	query := `INSERT INTO comments(email, comment)
			  VALUES(?, ?)`
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id: ", insertId)
}

// Prepare Statement
func TestPrepStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `INSERT INTO comments(email, comment)
			  VALUES(?, ?)`
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	// Exec Statement
	for i := 0; i < 10; i++ {
		email := "boy" + strconv.Itoa(i) + "@mail.com"
		comment := "boyya: " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id: ", insertId)
	}
}

// Transaction
func TestTransac(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transac
	query := `INSERT INTO comments(email, comment)
	VALUES(?, ?)`

	for i := 0; i < 10; i++ {
		email := "boy" + strconv.Itoa(i) + "@mail.com"
		comment := "boyya: " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		lastId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id: ", lastId)
	}

	// commit / success case
	// err = tx.Commit()

	// rollback / rollback case
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

}
