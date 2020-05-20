package handlers

import (
	"github.com/sbeverly/auth/internal/jwt"
	"github.com/sbeverly/auth/internal/db"
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

type AuthenticatedContext struct {
	echo.Context
}

func (c *AuthenticatedContext) GetUser() (*db.User, error) {
	cookie, err := c.Cookie("token")

	if err != nil {
		return nil, err
	}

	claims, _ := jwt.GetClaims(cookie.Value)

	conn := db.Start()
	user, err := conn.GetUserByID(claims.UserID)
	conn.End()

	if err != nil {
		return nil, err
	}

	return user, nil
}
