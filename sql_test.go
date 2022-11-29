package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// bisa diganti delete atau update
	query := "INSERT INTO customer(id, name) VALUES('pradana', 'Pradana')"

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// bisa diganti delete atau update
	query := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// bisa diganti delete atau update
	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("========================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)

		if email.Valid {
			fmt.Println("Email:", email.String)
		}

		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)

		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}

		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}
	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	// bisa diganti delete atau update
	query := "SELECT username FROM users WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	// bisa diganti delete atau update
	query := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "ferrian"
	password := "ferrian"

	// bisa diganti delete atau update
	query := "INSERT INTO users(username, password) VALUES(?, ?)"

	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}