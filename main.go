package main

import (
	"github.com/chrisbirster/webgamedev/pkg/logger"
	"github.com/chrisbirster/webgamedev/pkg/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(logger.NewLoggerWithConfig())

	// markdown content
	metadata, markdown, err := ParseMarkdownFile("content/emails-with-react.md")
	if err != nil {
		e.Logger.Errorf("Error parsing markdown file: %v", err)
		// Logger.LogAttrs(context.Background(), slog.LevelError, "MARKDOWN", slog.String("err", err.Error()))
	}
	e.Logger.Infof("Metadata: %v", metadata)
	e.Logger.Infof("Markdown: %v", markdown)
	// Logger.LogAttrs(context.Background(), slog.LevelInfo, "MARKDOWN", slog.Any("METADATA", metadata), slog.Any("CONTENT", markdown))

	renderer, err := NewTemplateRenderer()
	if err != nil {
		// Logger.LogAttrs(context.Background(), slog.LevelError, "TEMPLATE", slog.String("err", err.Error()))
		e.Logger.Errorf("Error creating template renderer: %v", err)
	}
	e.Renderer = renderer

	// static files
	e.Static("/css", "css")
	e.Static("/static", "static")

	// routes
	routes.HandleRoutes(e)

	// logger.LogAttrs(context.Background(), slog.LevelInfo, "SERVER", slog.String("PORT", ":42069"))
	e.Logger.Fatal(e.Start(":42069"))
}
