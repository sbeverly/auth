package main

import (
	"github.com/sbeverly/auth/db"
	"golang.org/x/crypto/bcrypt"
	"log"
)

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

func main() {
	createUser("Siyan Beverly", "siyan.beverly@gmail.com", "12345")
}
