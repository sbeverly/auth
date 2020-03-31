package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	STATUS_OK = "OK"
)

var startTime time.Time

func getUptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}

func Ping(c echo.Context) error {
	uptime := getUptime()
	return c.JSON(http.StatusOK, &PingResponse{STATUS_OK, uptime.String()})
}
