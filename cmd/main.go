package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"template_htmx/handlers"
)

func main() {
	app := echo.New()

	// Set the logger level to log all levels of messages
	app.Logger.SetLevel(log.INFO)

	rootHandler := handlers.RootHandler{}
	app.GET("/", rootHandler.HandleRootRoute)

	// Start the server on port 8080
	err := app.Start(":8080")
	if err != nil {
		// If there's an error starting the server, log the error and exit
		app.Logger.Fatal(fmt.Errorf("error starting server: %w", err))
	}
}
