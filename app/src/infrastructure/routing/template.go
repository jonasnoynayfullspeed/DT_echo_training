package routing

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// SetTemplate Setting Template(Html)
func SetTemplate(e *echo.Echo) {
	t := &Template{
		templates: template.Must(template.New("").Funcs(template.FuncMap{
			"safeHTML": func(content string) template.HTML {
				return template.HTML(content)
			},
		}).ParseGlob("template/*.html")),
	}
	e.Renderer = t
}
