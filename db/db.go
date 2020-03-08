package db

import (
	"context"
	"fmt"
	//"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sbeverly/auth/config"
	"log"
)

var dbConf config.DatabaseConfig

type UserAccount struct {
	ID       string
	Name     string
	Email    string
	Password string
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

func (c Conn) CreateUser(name string, email string, password string) error {
	sql := `INSERT INTO user_account (name, email, password) 
		VALUES($1, $2, $3)`

	_, err := c.Exec(context.Background(), sql, name, email, password)

	if err != nil {
		return err
	}

	return nil
}
