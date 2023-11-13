package main

import (
	"github.com/chrisbirster/webgamedev/pkg/logger"
	"github.com/chrisbirster/webgamedev/pkg/renderer"
	"github.com/chrisbirster/webgamedev/pkg/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(logger.NewLoggerWithConfig())

	e.Renderer = renderer.NewTemplateRenderer()

	// static files
	e.Static("/css", "css")
	e.Static("/static", "static")

	// routes
	routes.HandleRoutes(e)

	e.Logger.Fatal(e.Start(":42069"))
}
