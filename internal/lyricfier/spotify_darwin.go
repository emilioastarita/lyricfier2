// +build darwin
package lyricfier

import (
	"golang.org/x/text/encoding/charmap"
	"log"
	"os/exec"
	"strings"
	"time"
)

type Spotify struct {
}

func (h *Spotify) Init() {
}
func getDataFromAppleScript() (bool, string, string) {
	cmd := exec.Command("osascript", "-e", `
tell application "System Events"
   set processList to (name of every process)
end tell
if (processList contains "Spotify") is true then
   tell application "Spotify"
      set artistName to artist of current track
      set trackName to name of current track
      return trackName & " - " & artistName
    end tell
end if
`)

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return false, "", ""
	}
	d := charmap.CodePage850.NewDecoder()
	outDecoded, err := d.Bytes(out)
	s := string(outDecoded)
	parts := strings.Split(s, " - ")
	if len(parts) == 2 {
		return true, parts[0], parts[1]
	}
	return false, "", ""
}

func (h *Spotify) GetMetadata(newSong chan *Song) {
	metadata := &Song{}
	found, title, artist := getDataFromAppleScript()
	if !found {
		newSong <- nil
		return
	}
	metadata.Title = title
	metadata.Artist = artist
	newSong <- metadata
}

func (h *Spotify) Ticker(changes chan string) {
	fpsTicker := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-fpsTicker.C:
			ok, _, _ := getDataFromAppleScript()
			if ok {
				changes <- "yes"
			}
		}
	}
}

func GetDbPath() string {
	return filepath.join(os.Getenv("HOME"), "/Library/Application Support/lyricfier")
}

func GetPlatformName() string {
	return "darwin"
}
