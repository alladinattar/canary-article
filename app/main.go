package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	if os.Getenv("VERSION") == "v2" {
		i := 0
		e.GET("/", func(c echo.Context) error {
			if i%2 == 0 {
				i++
				return c.String(http.StatusOK, "You use v2 app version!!!")
			} else {
				i++
				return c.String(http.StatusInternalServerError, "App crashed")
			}
		})
	} else if os.Getenv("VERSION") == "v3" {
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "You use v3 app version!!!")
		})
	} else {
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "You use v1 app version!!!")
		})
	}
	e.Logger.Fatal(e.Start(":80"))
}
