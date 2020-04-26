package lyricfier

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"mime"
	"net/http"
	"os"
)

type Server struct {
	e       *echo.Echo
	appData *AppData
	hub     *Hub
}

type TemplateRegistry struct {
	templates *template.Template
}

type SongPostData struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Lyric  string `json:"lyric"`
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (h *Server) Init(appData *AppData) {
	mime.AddExtensionType(".mjs", "application/javascript")
	h.hub = newHub()
	go h.hub.run()
	h.e = echo.New()
	h.appData = appData
	h.routes(h.hub)
	s := &http.Server{
		Addr: appData.Address,
	}
	h.e.HidePort = true
	h.e.HideBanner = true
	go func() {
		h.e.Logger.Fatal(h.e.StartServer(s))
	}()
}

var useLocal = os.Getenv("LOCAL_ASSETS") == "true"

func (h *Server) routes(hub *Hub) {
	h.e.GET("/songs", func(c echo.Context) error {
		var list []*SongItem
		if err := ListSongs(&list); err != nil {
			fmt.Printf("list_songs %v\n", err)
			return c.JSON(http.StatusInternalServerError, h.appData)
		}
		return c.JSON(http.StatusOK, list)
	})
	h.e.POST("/save-song", func(c echo.Context) error {
		s := new(SongPostData)
		if err := c.Bind(s); err != nil {
			fmt.Printf("save_song %v\n", err)
			return c.JSON(http.StatusInternalServerError, h.appData)
		}
		key := SongKey(s.Artist, s.Title)
		err := Write(SongsBucket, key, []byte(s.Lyric))
		if err != nil {
			fmt.Printf("save_song %v\n", err)
			return c.JSON(http.StatusInternalServerError, h.appData)
		}
		if h.appData.Song.Title == s.Title && h.appData.Song.Artist == s.Artist {
			h.appData.Song.Lyric = s.Lyric
		}
		return c.JSON(http.StatusOK, h.appData)
	})
	h.e.POST("/save-settings", func(c echo.Context) error {
		s := new(Settings)
		if err := c.Bind(s); err != nil {
			fmt.Printf("save_song %v\n", err)
			return c.JSON(http.StatusInternalServerError, h.appData)
		}
		settings, err := json.Marshal(s)
		if err != nil {
			fmt.Printf("could not marshal config json: %v\n", err)
			return c.JSON(http.StatusInternalServerError, h.appData)
		}

		err = Write(GeneralBucket, SettingsKey, settings)
		if err != nil {
			fmt.Printf("save_song: %v", err)
			return c.JSON(http.StatusInternalServerError, h.appData)
		}
		h.appData.Settings = s
		return c.JSON(http.StatusOK, h.appData)
	})

	fs := http.FileServer(FS(useLocal))

	h.e.GET("/*", echo.WrapHandler(fs))

	h.e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, FSMustString(useLocal, "/static/index.html"))
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
