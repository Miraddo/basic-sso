package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/miraddo/basic-sso/app/tooling/token"
)

type User struct {
	Username string
	Password string
}

var connect *sql.DB

func init() {
	conn, err := sql.Open("mysql", "root:root@tcp(dbContainer)/sso")

	if err != nil {
		log.Println(err)
	}
	connect = conn

}

func (u *User) Insert() (bool, error) {
	t, err := token.GenToken()

	defer connect.Close()

	if err != nil {
		log.Printf("got an error wile generate token, %v", err)
		return false, err
	}

	stmtIns, err := connect.Prepare("INSERT INTO users (username, password, token) VALUES (?,?,?)")

	if err != nil {
		return false, err
	}

	defer stmtIns.Close()

	_, err = stmtIns.Exec(u.Username, u.Password, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
