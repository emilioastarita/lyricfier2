// +build linux
package lyricfier

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"os"
)

type Spotify struct {
	conn *dbus.Conn
	bus  dbus.BusObject
	ch   chan *dbus.Signal
}

func (h *Spotify) Init() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	h.conn = conn
	h.bus = conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2")
	h.bus.AddMatchSignal("org.freedesktop.DBus.Properties", "PropertiesChanged", dbus.WithMatchObjectPath(h.bus.Path()))
	h.ch = make(chan *dbus.Signal)
	h.conn.Signal(h.ch)
}

func (h *Spotify) GetMetadata(newSong chan *Song) {
	metadata := &Song{}
	res, err := h.bus.GetProperty("org.mpris.MediaPlayer2.Player.Metadata")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get metadata from Spotify", err)
		newSong <- nil
		return
	}
	m, ok := res.Value().(map[string]dbus.Variant)
	if ok {
		artists := m["xesam:artist"].Value().([]string)
		metadata.ArtUrl = m["mpris:artUrl"].Value().(string)
		metadata.Artist = ""
		if len(artists) > 0 {
			metadata.Artist = artists[0]
		}
		metadata.Title = m["xesam:title"].Value().(string)
		newSong <- metadata
	}
}

func (h *Spotify) Ticker(changes chan string) {
	for {
		select {
		case <-h.ch:
			changes <- "yes"
		}
	}
}
