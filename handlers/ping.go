package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	STATUS_OK = "OK"
)

func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, &PingResponse{STATUS_OK})
}
