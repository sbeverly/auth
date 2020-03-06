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
	err := db.CreateUser(name, email, string(hPwd))

	if err != nil {
		log.Println(err)
	}
}

func main() {
	createUser("Siyan Beverly", "siyan.beverly@gmail.com", "12345")
}
