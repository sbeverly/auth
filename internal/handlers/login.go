package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sbeverly/auth/internal/db"
	"github.com/sbeverly/auth/internal/jwt"
	"golang.org/x/crypto/bcrypt"
	"github.com/sbeverly/auth/internal/config"
	"net/http"
	"time"
)

const (
	invalidCredsMSG      = "Wrong Username/Password"
	noEmailPasswordMSG 	 = "Could not extract email/password from request"
)

var cookieConf config.CookieConfig

func init() {
	cookieConf = config.GetConfig().Cookie
}

// Login : Handle login request
func Login(c echo.Context) error {
	req := new(LoginRequest)

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

	token, err := jwt.Generate(&jwt.Payload{Email: user.Email})

	if err != nil {
		c.NoContent(http.StatusInternalServerError)
	}

	cookie := makeCookie(token)
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// UTILS

func makeCookie(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "auth"
	cookie.Value = token
	cookie.Domain = cookieConf.Domain
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Expires = time.Now().Add(24 * time.Hour)
	return cookie
}
