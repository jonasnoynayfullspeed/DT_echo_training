package routing

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"time"
)

type Template struct {
	templates *template.Template
}

const ReadableDateFormat = "January 02, 2006"

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// SetTemplate Setting Template(Html)
func SetTemplate(e *echo.Echo) {
	t := &Template{
		templates: template.Must(template.New("").Funcs(template.FuncMap{
			"toHTML": func(content string) template.HTML {
				return template.HTML(content)
			},
			"readableDateFormat": func(time time.Time) string {
				return time.Format(ReadableDateFormat)
			},
		}).ParseGlob("template/*.html")),
	}
	e.Renderer = t
}
