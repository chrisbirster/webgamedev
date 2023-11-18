package routes

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func HandleRoutes(e *echo.Echo, db *sql.DB) {

	e.GET("/", func(c echo.Context) error {
		return HandleIndex(c, db)
	})

	e.GET("/blog", func(c echo.Context) error {
		return HandleBlogIndex(c, db)
	})

	e.GET("/deez", func(c echo.Context) error {
		return HandleDeez(c, db)
	})

	e.GET("/arcade", func(c echo.Context) error {
		return HandleArcade(c)
	})

	e.GET("/blog/:slug", func(c echo.Context) error {
		var slug string = c.Param("slug")
		return HandleBlogPost(c, db, slug)
	})
}
