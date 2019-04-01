package lyricfier

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

type Server struct {
	e       *echo.Echo
	appData *AppData
	hub     *Hub
}

type TemplateRegistry struct {
	templates *template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (h *Server) Init(appData *AppData) {
	h.hub = newHub()
	go h.hub.run()
	h.e = echo.New()
	h.appData = appData
	h.e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	h.routes(h.hub)
	h.e.Logger.Fatal(h.e.Start(appData.Address))
}

func (h *Server) routes(hub *Hub) {
	h.e.Static("static", "static")
	h.e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", h.appData)
	})
	h.e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, h.appData)
	})
	h.e.GET("/ws", func(c echo.Context) error {
		return serveWs(hub, c)
	})
}

func (h *Server) NotifyChanges() {
	h.hub.broadcast <- []byte("")
}
