package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"monkey/core"
	"monkey/crud"
	"net/http"
	"os"
)

func run() error {
	// Environmental Variables

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${remote_ip}<${method}>\t\t status=${status}\t\t ↘=${bytes_in}b ↗=${bytes_out}b\t\t latency=${latency_human}\t uri=${uri}\t\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	// Test Database

	test := crud.IsDatabaseWorking()
	if test != nil {
		panic(test)
	}

	// Initialize the Database
	crud.InitializeUsers()

	MountApiRouter(e)

	static := e.Group("/")
	static.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "static/",
		HTML5: true,
	}))

	e.Logger.Info(e.Start(core.Settings.ServerPort))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
