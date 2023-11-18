package routes

import (
	"database/sql"
	"net/http"
	"path/filepath"

	"github.com/chrisbirster/webgamedev/pkg/markdown"
	"github.com/chrisbirster/webgamedev/pkg/models"

	"github.com/labstack/echo/v4"
)

func HandleBlogIndex(c echo.Context, db *sql.DB) error {
	contentDirectory := filepath.Join("content")
	files, err := filepath.Glob(filepath.Join(contentDirectory, "*.md"))
	if err != nil {
		// Handle error (e.g., directory not found)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading content directory")
	}

	var posts []models.Post
	for _, file := range files {
		ligma, err := markdown.ParseMarkdownFile(file)
		if err != nil {
			// Log the error but continue processing other files
			c.Logger().Errorf("Error parsing markdown file: %s, %v", file, err)
			continue
		}

		post := models.Post{
			Title:       ligma.Metadata.Title,
			Description: "provide a better excerpt here",
			MetaData: models.MetaData{
				Arthor: models.Arthor{
					Name: ligma.Metadata.Author,
				},
				Date:       ligma.Metadata.Published,
				Categories: ligma.Metadata.Category,
				Tags:       ligma.Metadata.Tags,
			},
		}
		posts = append(posts, post)
	}

	data := models.Page{
		Name:        "<WGD/>",
		MainHeading: "Welcome to my website!",
		SubHeading:  "This is a subheader",
		Title:       "WebGameDev",
		NavLinks: []models.Link{
			{Name: "Courses", Path: "/", IsCurrentPage: true},
			{Name: "Books", Path: "/about", IsCurrentPage: false},
			{Name: "Arcade", Path: "/contact", IsCurrentPage: false},
		},
		Posts: posts,
	}

	c.Logger().Infof("about to call c.Render with DATA: %v", data)
	return c.Render(http.StatusOK, "blog", data)
}
