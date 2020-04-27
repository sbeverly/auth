package handlers

import (
	"net/http"

	"github.com/sbeverly/auth/internal/cookies"
	"github.com/labstack/echo/v4"
)

// Logout : /api/logout
func Logout(c echo.Context) error {
	ac := c.(*AuthenticatedContext)

	ck := cookies.GenerateLogoutCookie()
	ac.SetCookie(ck)
	return ac.NoContent(http.StatusOK)
}
