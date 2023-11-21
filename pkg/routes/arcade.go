package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleArcade(c echo.Context) error {
	fmt.Println("arcade route called")
	return c.Render(http.StatusOK, "arcade", nil)
}
