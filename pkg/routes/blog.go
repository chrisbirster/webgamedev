package routes

import (
	"database/sql"
	"net/http"

	"github.com/chrisbirster/webgamedev/pkg/models"

	"github.com/labstack/echo/v4"
)

func HandleBlogIndex(c echo.Context, db *sql.DB) error {
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
		Posts: []models.Post{
			{
				Title:       "Extend Vite's Config. Customize Your Development Experience.",
				Description: "Delve into the world of Vite and unlock the full potential of your development server. Learn how to take control of Vite's powerful features to streamline your workflow, enhance performance, and tailor the server to your project's specific needs. Perfect for developers looking to elevate their Vite experience!",
				MetaData: models.MetaData{
					Arthor: models.Arthor{
						Name: "GM",
						Link: "https://www.google.com",
					},
					Date:       "2023-11-11",
					Categories: []string{"Front End", "Development"},
					Tags:       []string{"JavaScript", "Vite", "Configuration"},
				},
			},
			{
				Title:       "Extend Vite's Config. Customize Your Development Experience.",
				Description: "Delve into the world of Vite and unlock the full potential of your development server. Learn how to take control of Vite's powerful features to streamline your workflow, enhance performance, and tailor the server to your project's specific needs. Perfect for developers looking to elevate their Vite experience!",
				MetaData: models.MetaData{
					Arthor: models.Arthor{
						Name: "GM",
						Link: "https://www.google.com",
					},
					Date:       "2023-11-11",
					Categories: []string{"Front End", "Development"},
					Tags:       []string{"JavaScript", "Vite", "Configuration"},
				},
			},
			{
				Title:       "Extend Vite's Config. Customize Your Development Experience.",
				Description: "Delve into the world of Vite and unlock the full potential of your development server. Learn how to take control of Vite's powerful features to streamline your workflow, enhance performance, and tailor the server to your project's specific needs. Perfect for developers looking to elevate their Vite experience!",
				MetaData: models.MetaData{
					Arthor: models.Arthor{
						Name: "GM",
						Link: "https://www.google.com",
					},
					Date:       "2023-11-11",
					Categories: []string{"Front End", "Development"},
					Tags:       []string{"JavaScript", "Vite", "Configuration"},
				},
			},
			{
				Title:       "Extend Vite's Config. Customize Your Development Experience.",
				Description: "Delve into the world of Vite and unlock the full potential of your development server. Learn how to take control of Vite's powerful features to streamline your workflow, enhance performance, and tailor the server to your project's specific needs. Perfect for developers looking to elevate their Vite experience!",
				MetaData: models.MetaData{
					Arthor: models.Arthor{
						Name: "GM",
						Link: "https://www.google.com",
					},
					Date:       "2023-11-11",
					Categories: []string{"Front End", "Development"},
					Tags:       []string{"JavaScript", "Vite", "Configuration"},
				},
			},
			{
				Title:       "Extend Vite's Config. Customize Your Development Experience.",
				Description: "Delve into the world of Vite and unlock the full potential of your development server. Learn how to take control of Vite's powerful features to streamline your workflow, enhance performance, and tailor the server to your project's specific needs. Perfect for developers looking to elevate their Vite experience!",
				MetaData: models.MetaData{
					Arthor: models.Arthor{
						Name: "GM",
						Link: "https://www.google.com",
					},
					Date:       "2023-11-11",
					Categories: []string{"Front End", "Development"},
					Tags:       []string{"JavaScript", "Vite", "Configuration"},
				},
			},
		},
	}
	c.Logger().Infof("about to call c.Render with DATA: %v", data)
	return c.Render(http.StatusOK, "blog", data)
}
