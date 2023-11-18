package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleArcade(c echo.Context) error {
	return c.Render(http.StatusOK, "arcade", nil)
}
