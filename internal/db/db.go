package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/sbeverly/auth/internal/config"
)

var dbConf config.DatabaseConfig

type User struct {
	Name    string
	Email   string
	IsAdmin bool
}

func init() {
	dbConf = config.GetConfig().DB
}

type Conn struct {
	*pgx.Conn
}

func Start() *Conn {
	connStr := fmt.Sprintf("user=%s password=%s host=/cloudsql/%s database=authentication",
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

func (c Conn) GetUserByEmail(email string) (*User, error) {
	var name string
	var isAdmin bool

	sql := `SELECT name, is_admin FROM user_account WHERE email = $1`
	err := c.QueryRow(context.Background(), sql, email).Scan(&name, &isAdmin)

	if err != nil {
		return nil, err
	}

	usr := &User{name, email, isAdmin}

	return usr, nil
}

func (c Conn) GetUserWithPassword(email string) (*User, string, error) {
	var name string
	var encryptedPassword string
	var isAdmin bool

	sql := `SELECT name, password, is_admin FROM user_account WHERE email = $1`
	err := c.QueryRow(context.Background(), sql, email).Scan(&name, &encryptedPassword, &isAdmin)

	if err != nil {
		return nil, "", err
	}

	return &User{name, email, isAdmin}, encryptedPassword, nil
}

func (c Conn) CreateUser(name string, email string, password string) error {
	sql := `INSERT INTO user_account (name, email, password) VALUES($1, $2, $3)`

	_, err := c.Exec(context.Background(), sql, name, email, password)

	if err != nil {
		return err
	}

	return nil
}
