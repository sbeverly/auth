package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MeResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Me : /me request handler
func Me(c echo.Context) error {
	ac := c.(*AuthenticatedContext)
	user, _ := ac.GetUser()

	return ac.JSON(http.StatusOK, &MeResponse{user.Name, user.Email})
}
