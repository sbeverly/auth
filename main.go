package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sbeverly/auth/internal/handlers"
	"os"
)

/*
const (
	PWDCOST = 10
)

func createUser(name string, email string, pwd string) {
	hPwd, _ := bcrypt.GenerateFromPassword([]byte(pwd), PWDCOST)
	conn := db.Start()
	err := conn.CreateUser(name, email, string(hPwd))
	conn.End()
	if err != nil {
		log.Println(err)
	}
}
*/

func main() {
	port := "1323"
	if val := os.Getenv("PORT"); val != "" {
		port = val
	}

	e := echo.New()
	e.GET("/api/ping", handlers.Ping)
	e.POST("/api/login", handlers.Login)
	e.POST("/api/verify", handlers.Verify)
	e.Logger.Fatal(e.Start(":" + port))
}
