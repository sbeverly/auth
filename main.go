package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sbeverly/auth/internal/handlers"
	"github.com/sbeverly/auth/internal/jwt"
	"net/http"
	"os"
)

func validateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")

		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		err = jwt.Verify(cookie.Value)

		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		return next(&handlers.AuthenticatedContext{c})
	}
}

func adminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ac := c.(*handlers.AuthenticatedContext)

		user, err := ac.GetUser()

		if user.IsAdmin != true {
			return c.NoContent(http.StatusUnauthorized)
		}

		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return next(&handlers.AuthenticatedContext{c})
	}
}

func main() {
	e := echo.New()

	e.Pre(middleware.HTTPSRedirect())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://login-local.siyan.io", "https://login.siyan.io"},
		AllowCredentials: true,
	}))

	e.GET("/ping", handlers.Ping)
	e.POST("/login", handlers.Login)
	
	user := e.Group("/user", validateToken)
	user.GET("/verify", handlers.Verify)
	user.GET("/me", handlers.Me)
	user.POST("/logout", handlers.Logout)
	
	admin := e.Group("/admin", validateToken, adminOnly)
	admin.POST("/user/create", handlers.CreateUser)

	port := "1323"
	if val := os.Getenv("PORT"); val != "" {
		port = val
	}

	e.Logger.Fatal(e.Start(":" + port))
}
