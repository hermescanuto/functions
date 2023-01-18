package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func get_port() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"msg": "Hello World"})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/api/hello", hello)
	e.Logger.Fatal(e.Start(get_port()))
}
