package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sbeverly/auth/config"
)

var dbConf config.DatabaseConfig

// NOTE: May move to /models
type UserAccount struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func init() {
	dbConf = config.GetConfig().DB
}

func Connect() *pgx.Conn {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/authentication",
		dbConf.User,
		dbConf.Password,
		dbConf.Host)
	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		panic(err)
	}
	return conn
}

func Exec(sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	conn := Connect()
	defer conn.Close(context.Background())

	return conn.Exec(context.Background(), sql, arguments...)
}

func CreateUser(name string, email string, password string) error {
	sql := `INSERT INTO user_account (name, email, password) 
		VALUES($1, $2, $3)`

	_, err := Exec(sql, name, email, password)

	if err != nil {
		return err
	}

	return nil
}
