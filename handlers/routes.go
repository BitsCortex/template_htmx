package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type RootHandler struct{}

func (h RootHandler) HandleRootRoute(c echo.Context) error {
	// If the server starts successfully, print a success message
	w := "Gratz app started, so have fun!!!"
	n := time.Now()
	fmt.Println("[", n.Format("2006-01-02 15:04:05"), "]", w)

	return c.String(http.StatusOK, "Hello, World!")
}
