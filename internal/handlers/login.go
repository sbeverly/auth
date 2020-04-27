package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sbeverly/auth/internal/db"
	"github.com/sbeverly/auth/internal/jwt"
	"github.com/sbeverly/auth/internal/cookies"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const (
	invalidCredsMSG      = "Wrong Username/Password"
	noEmailPasswordMSG 	 = "Could not extract email/password from request"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login : /login
func Login(c echo.Context) error {
	req := new(loginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{noEmailPasswordMSG})
	}

	conn := db.Start()
	user, encryptedPassword, err := conn.GetUserWithPassword(req.Email)
	conn.End()

	if err != nil {
		// user not found
		return c.JSON(http.StatusUnauthorized, &ErrorResponse{invalidCredsMSG})
	}

	err = bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(req.Password))

	if err != nil {
		return c.JSON(http.StatusUnauthorized, &ErrorResponse{invalidCredsMSG})
	}

	token, err := jwt.Generate(&jwt.Payload{UserID: user.ID})

	if err != nil {
		c.NoContent(http.StatusInternalServerError)
	}

	ck := cookies.GenerateLoginCookie(token)
	c.SetCookie(ck)
	return c.NoContent(http.StatusOK)
}

