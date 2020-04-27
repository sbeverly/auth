package handlers

import (
	"github.com/sbeverly/auth/internal/jwt"
	"github.com/sbeverly/auth/internal/db"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type VerifyRequest struct {
	Token string `json:"token"`
}

type MeRequest struct{}

type MeResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type SuccessResponse struct {
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

type PingResponse struct {
	Status string `json:"status, omitempty"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password  string `json:"password"`
}

type AuthenticatedContext struct {
	echo.Context
}

func (c *AuthenticatedContext) GetUser() (*db.User, error) {
	cookie, err := c.Cookie("token")

	if err != nil {
		return nil, err
	}

	payload, _ := jwt.Claims(cookie.Value)

	conn := db.Start()
	user, err := conn.GetUserByID(payload.UserID)
	conn.End()

	if err != nil {
		return nil, err
	}

	return user, nil
}
