package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sbeverly/auth/db"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

const (
	PWDCOST = 10
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func createUser(name string, email string, pwd string) {
	hPwd, _ := bcrypt.GenerateFromPassword([]byte(pwd), PWDCOST)
	conn := db.Start()
	err := conn.CreateUser(name, email, string(hPwd))
	conn.End()
	if err != nil {
		log.Println(err)
	}
}

func login(c echo.Context) error {
	req := new(LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Could not extract email/password from request.")
	}

	conn := db.Start()
	user, pwdHash, err := conn.GetUser(req.Email)
	conn.End()

	if err != nil {
		// user not found
		return c.JSON(http.StatusUnauthorized, "Invalid Credentials.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(req.Password))

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials.")
	}

	return c.JSON(http.StatusOK, &LoginResponse{
		Token: "GoodLogin",
		Email: user.Email})
}

func main() {
	e := echo.New()
	e.POST("/login", login)
	e.Logger.Fatal(e.Start(":1323"))
}
