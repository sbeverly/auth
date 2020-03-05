package main

import (
	"github.com/sbeverly/auth/config"
	"github.com/sbeverly/auth/db"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const (
	PWDCOST = 10
)

func createUser(email string, pwd string) {
	hPwd, _ := bcrypt.GenerateFromPassword([]byte(pwd), PWDCOST)
	log.Println(hPwd)
}

func main() {
	conf := config.GetConfig()
	db.Connect(conf)
	createUser("siyan.beverly", "12345")
}
