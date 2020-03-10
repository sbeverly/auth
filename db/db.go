package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/sbeverly/auth/config"
	"log"
)

var dbConf config.DatabaseConfig

type User struct {
	Name  string
	Email string
}

func init() {
	dbConf = config.GetConfig().DB
}

type Conn struct {
	*pgx.Conn
}

func Start() *Conn {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/authentication",
		dbConf.User,
		dbConf.Password,
		dbConf.Host)
	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		log.Fatal(err)
	}
	return &Conn{conn}
}

func (c Conn) End() {
	c.Close(context.Background())
}

func (c Conn) GetUser(email string) (*User, string, error) {
	var name string
	var pwdHash string

	sql := `SELECT name, password
		FROM user_account
		WHERE email = $1`
	err := c.QueryRow(context.Background(), sql, email).Scan(&name, &pwdHash)

	if err != nil {
		return nil, "", err
	}

	return &User{
		name,
		email}, pwdHash, nil
}

func (c Conn) CreateUser(name string, email string, password string) error {
	sql := `INSERT INTO user_account (name, email, password) 
		VALUES($1, $2, $3)`

	_, err := c.Exec(context.Background(), sql, name, email, password)

	if err != nil {
		return err
	}

	return nil
}
