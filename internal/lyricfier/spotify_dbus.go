package lyricfier

import (
	"fmt"
	"github.com/godbus/dbus"
	"os"
)

type SpotifyDbus struct {
	conn    *dbus.Conn
	bus     dbus.BusObject
	Changes chan *dbus.Signal
}

func (h *SpotifyDbus) Init() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	h.conn = conn
	h.bus = conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2")
	h.bus.AddMatchSignal("org.freedesktop.DBus.Properties", "PropertiesChanged", dbus.WithMatchObjectPath(h.bus.Path()))
	h.Changes = make(chan *dbus.Signal)
	h.conn.Signal(h.Changes)
}

func (h *SpotifyDbus) GetMetadata(newSong chan *Song) {
	metadata := &Song{}
	res, err := h.bus.GetProperty("org.mpris.MediaPlayer2.Player.Metadata")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get metadata from Spotify", err)
		newSong <- nil
		return
	}
	m, ok := res.Value().(map[string]dbus.Variant)

	if ok {
		metadata.ArtUrl = m["mpris:artUrl"].Value().(string)
		metadata.Artist = m["xesam:artist"].Value().([]string)[0]
		metadata.Title = m["xesam:title"].Value().(string)
		newSong <- metadata
	}

}
