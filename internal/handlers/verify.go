package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sbeverly/auth/internal/jwt"
	"net/http"
)

const (
	UNAUTHORIZED     = "Unauthorized"
	BAD_REQ_NO_TOKEN = "Token is missing or malformed"
)

func Verify(c echo.Context) error {
	req := new(VerifyRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorResponse{BAD_REQ_NO_TOKEN})
	}

	err := jwt.Verify(req.Token)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, &ErrorResponse{UNAUTHORIZED})
	}

	return c.JSON(http.StatusOK, &SuccessResponse{})
}
