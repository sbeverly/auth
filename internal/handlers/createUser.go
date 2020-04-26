package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sbeverly/auth/internal/db"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser : /user/create handler
func CreateUser(c echo.Context) error {
	req := new(CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	conn := db.Start()
	err := conn.CreateUser(req.Name, req.Email, string(hashedPass))
	conn.End()

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
