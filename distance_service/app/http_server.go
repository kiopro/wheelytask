package app

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// StartHTTPServer : start HTTP server
func StartHTTPServer() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Wheely Distance Info Service v0.1")
	})

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "I'm alive!")
	})

	HTTPPort := os.Getenv("HTTP_PORT")

	e.Logger.Fatal(e.Start(HTTPPort))
}
