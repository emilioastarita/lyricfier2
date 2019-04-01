package lyricfier

import (
	"html/template"
	"io"
	"net/http"
	"github.com/labstack/echo"
)

type Server struct {
	e *echo.Echo
	appData *AppData
}

type TemplateRegistry struct {
	templates *template.Template
}
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func (h *Server) Init(appData *AppData) {
	h.e = echo.New()
	h.appData = appData
	h.e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	h.routes()
	h.e.Logger.Fatal(h.e.Start(":1323"))
}

func (h *Server) routes() {
	h.e.Static("/public/static", "static")
	h.e.GET("/status", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.html", h.appData)
	})
	h.e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, h.appData)
	})
}

