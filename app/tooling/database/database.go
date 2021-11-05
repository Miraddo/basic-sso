// Package database is for handle connection with MariaDB.
package database

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/miraddo/basic-sso/app/tooling/hash"
)

// User structure.
type User struct {
	Username string
	Password string
}

// init will be open a database connection.
func new() *sql.DB {
	conn, err := sql.Open("mysql", "root:root@tcp(dbContainer)/sso")

	if err != nil {
		log.Fatal("Failed to init db:", err)
	}

	return conn
}

// check will be check for the username is exist or not.
func (u *User) check(db *sql.DB) bool {

	var nRow int

	// Query to check that the username is exist or not.
	err := db.QueryRow(`SELECT count(id) 
								FROM users 
								WHERE username = ?`, u.Username).Scan(&nRow)

	switch {
	case err != nil:
		log.Println(err)
		return false

	case nRow > 0:
		return true

	default:
		return false
	}
}

// Insert will be insert new user to the database.
func (u *User) Insert() (bool, error) {

	db := new()

	if !u.check(db) {
		return false, errors.New("User [" + u.Username + "] already exist!")
	}

	// GenToken it will return a random token.
	t, err := hash.GenToken()

	if err != nil {
		log.Printf("got an error while generate token, %v", err)
		return false, err
	}

	// Query to insert new user to the database.
	stmtIns, err := db.Prepare("INSERT INTO users (username, password, token) VALUES (?,?,?)")

	if err != nil {
		return false, err
	}

	defer stmtIns.Close()

	// Execute the insert query with hashing the password.
	_, err = stmtIns.Exec(u.Username, hash.GenPass(u.Password), t)

	defer db.Close()

	if err != nil {
		return false, err
	}

	return true, nil
}
