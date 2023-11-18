package routes

import (
	"database/sql"
	"net/http"

	"github.com/chrisbirster/webgamedev/pkg/markdown"
	"github.com/chrisbirster/webgamedev/pkg/models"

	"github.com/labstack/echo/v4"
)

func HandleIndex(c echo.Context, db *sql.DB) error {
	// markdown content
	ligma, err := markdown.ParseMarkdownFile("content/emails-with-react.md")
	if err != nil {
		c.Logger().Errorf("Error parsing markdown file: %v", err)
	}
	c.Logger().Infof("Metadata: %v", ligma.Metadata)
	c.Logger().Infof("Markdown: %v", ligma.Markdown)

	data := models.Page{
		Name:        "<WGD/>",
		MainHeading: "Welcome to my website!",
		SubHeading:  "This is a subheader",
		Title:       "WebGameDev",
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
				Classes: "border border-blue-500",
			},
		},
		NavLinks: []models.Link{
			{Name: "Courses", Path: "/courses", IsCurrentPage: true},
			{Name: "Books", Path: "/books", IsCurrentPage: false},
			{Name: "Arcade", Path: "/arcade", IsCurrentPage: false},
			{Name: "Blog", Path: "/blog", IsCurrentPage: false},
		},
		Tutorials: []models.Tutorial{
			{Name: "NextJS", Link: "/tutorial/nextjs", Image: "nextjs-icon"},
			{Name: "React", Link: "/tutorial/react", Image: "react-icon"},
			{Name: "Astro", Link: "/tutorial/astro", Image: "astrojs-icon"},
		},
		Community: []models.Community{
			{Name: "@webgamedotdev", Link: "https://www.twitter.com/@webgamedotdev", Image: "xwitter-icon"},
			{Name: "@webgamedotdev", Link: "https://wwww.youtube.com/@webgamedotdev", Image: "youtube-icon"},
			{Name: "@webgamedotdev", Link: "https://discord.gg/z83Hpc8fEY", Image: "discord-icon"},
			{Name: "@webgamedotdev", Link: "https://github.com/webgamedotdev", Image: "github-icon"},
		},
		Course: models.Course{
			Name:  "Vite Custom Config Plugin",
			Link:  "/course/htmx-from-scratch",
			Image: "/static/youtube-thumbnail-vite-config-plugin.svg",
		},
		RecentPosts: []models.Post{
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
	return c.Render(http.StatusOK, "home", data)
}
