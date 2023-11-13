package template

import (
	"encoding/json"
	"html/template"
	"io"

	"github.com/Masterminds/sprig/v3"
	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	tmpl *template.Template
}

func NewTemplateRenderer() (*TemplateRenderer, error) {
	tmpl := template.New("").Funcs(sprig.FuncMap())
	_, err := tmpl.ParseGlob("templates/*.html")
	if err != nil {
		return nil, err
	}

	// Parse templates in the 'components' subdirectory.
	_, err = tmpl.ParseGlob("templates/components/*.html")
	if err != nil {
		return nil, err
	}

	// Parse templates in the 'layouts' subdirectory.
	_, err = tmpl.ParseGlob("templates/layouts/*.html")
	if err != nil {
		return nil, err
	}

	// Parse templates in the 'pages' subdirectory.
	_, err = tmpl.ParseGlob("templates/pages/*.html")
	if err != nil {
		return nil, err
	}

	return &TemplateRenderer{tmpl}, nil
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	dataBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		c.Logger().Errorf("Failed to marshal template data: %v", err)
	} else {
		c.Logger().Infof("Rendering template '%s' with data: %s", name, string(dataBytes))
	}
	return t.tmpl.ExecuteTemplate(w, "main.html", data)
}
