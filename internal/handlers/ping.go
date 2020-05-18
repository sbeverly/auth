package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PingResponse struct {
	Status string `json:"status, omitempty"`
}

func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, &PingResponse{"OK"})
}
