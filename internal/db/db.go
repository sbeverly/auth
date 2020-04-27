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
	ID      int
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
	connStr := fmt.Sprintf("user=%s password=%s host=/Users/siyanbeverly/Code/cloudsql/%s database=authentication",
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

func (c Conn) GetUserByID(userID int) (*User, error) {
	var ID int
	var name string
	var email string
	var isAdmin bool

	sql := `SELECT id, name, email, is_admin FROM user_account WHERE id = $1`
	err := c.QueryRow(context.Background(), sql, userID).Scan(&ID, &name, &email, &isAdmin)

	if err != nil {
		return nil, err
	}

	usr := &User{ID, name, email, isAdmin}

	return usr, nil
}

func (c Conn) GetUserByEmail(email string) (*User, error) {
	var ID int
	var name string
	var isAdmin bool

	sql := `SELECT id, name, is_admin FROM user_account WHERE email = $1`
	err := c.QueryRow(context.Background(), sql, email).Scan(&ID, &name, &isAdmin)

	if err != nil {
		return nil, err
	}

	usr := &User{ID, name, email, isAdmin}

	return usr, nil
}

func (c Conn) GetUserWithPassword(email string) (*User, string, error) {
	var ID int
	var name string
	var encryptedPassword string
	var isAdmin bool

	sql := `SELECT id, name, password, is_admin FROM user_account WHERE email = $1`
	err := c.QueryRow(context.Background(), sql, email).Scan(&ID, &name, &encryptedPassword, &isAdmin)

	if err != nil {
		return nil, "", err
	}

	return &User{ID, name, email, isAdmin}, encryptedPassword, nil
}

func (c Conn) CreateUser(name string, email string, password string) error {
	sql := `INSERT INTO user_account (name, email, password) VALUES($1, $2, $3)`

	_, err := c.Exec(context.Background(), sql, name, email, password)

	if err != nil {
		return err
	}

	return nil
}
