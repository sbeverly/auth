package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Verify(c echo.Context) error {
	// TODO: Validate current access using in memory token/user store
	return c.NoContent(http.StatusOK)
}
