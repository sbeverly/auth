package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sbeverly/auth/internal/db"
	"github.com/sbeverly/auth/internal/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

const (
	INVALID_CREDS_MSG      = "Wrong Username/Password"
	BAD_REQ_EMAIL_PASSWORD = "Could not extract email/password from request"
)

func makeCookie(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "auth"
	cookie.Value = token
	cookie.Secure = true
	cookie.Expires = time.Now().Add(24 * time.Hour)
	return cookie
}

func Login(c echo.Context) error {
	req := new(LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{BAD_REQ_EMAIL_PASSWORD})
	}

	conn := db.Start()
	_, pwdHash, err := conn.GetUser(req.Email)
	conn.End()

	if err != nil {
		// user not found
		return c.JSON(http.StatusUnauthorized, &ErrorResponse{INVALID_CREDS_MSG})
	}

	err = bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(req.Password))

	if err != nil {
		return c.JSON(http.StatusUnauthorized, &ErrorResponse{INVALID_CREDS_MSG})
	}

	token, err := jwt.Generate([]byte(`{"iss": "auth-server"}`))

	if err != nil {
		log.Println(err)
	}

	cookie := makeCookie(token)
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, &SuccessResponse{})
}
