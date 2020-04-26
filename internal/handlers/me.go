package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Me : /me request handler
func Me(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	user, _ := ac.GetUser()

	return c.JSON(http.StatusOK, &MeResponse{user.Name, user.Email})
}
