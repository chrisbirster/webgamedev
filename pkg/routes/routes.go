package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return HandleIndex(c)
	})

	e.GET("/blog", func(c echo.Context) error {
		return HandleBlogIndex(c)
	})

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test successful")
	})

	// e.GET("/blog/:slug", func(c echo.Context) error {
	// 	return HandleBlogPost(c)
	// })
}
