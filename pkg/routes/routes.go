package routes

import (
	"github.com/labstack/echo/v4"
)

func HandleRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return HandleIndex(c)
	})

	e.GET("/blog", func(c echo.Context) error {
		return HandleBlogIndex(c)
	})

	// e.GET("/blog/:slug", func(c echo.Context) error {
	// 	return HandleBlogPost(c)
	// })
}
