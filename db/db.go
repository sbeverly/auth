package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/sbeverly/auth/config"
)

func Connect(conf config.Config) *pgx.Conn {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/authentication",
		conf.DB.User,
		conf.DB.Password,
		conf.DB.Host)
	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		panic(err)
	}
	return conn
}
