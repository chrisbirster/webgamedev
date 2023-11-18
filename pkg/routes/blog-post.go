package routes

import (
	"database/sql"
	"net/http"
	"path/filepath"

	"github.com/chrisbirster/webgamedev/pkg/models"

	"github.com/labstack/echo/v4"

	"github.com/chrisbirster/webgamedev/pkg/markdown"
)

func HandleBlogPost(c echo.Context, db *sql.DB, slug string) error {
	// Construct the file path
	filePath := filepath.Join("content", slug+".md")

	// Use Ligma struct to parse the markdown file
	ligma, err := markdown.ParseMarkdownFile(filePath)
	if err != nil {
		// Handle error (file not found, etc.)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error loading blog post")
	}

	// Create a Post from the parsed markdown
	post := models.Post{
		Title:       ligma.Metadata.Title,
		Description: "TODO create an excerpt from the content",
		MetaData: models.MetaData{
			Arthor: models.Arthor{
				Name: ligma.Metadata.Author,
			},
		},
	}

	data := models.Page{
		Posts: []models.Post{post},
	}

	return c.Render(http.StatusOK, "blog-post", data)
}
