// +build windows
package lyricfier

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"log"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type Spotify struct {
}

func (h *Spotify) Init() {
}

func getDataFromTasklist() (bool, string, string) {
	cmd := exec.Command("tasklist.exe", "/fo", "csv", "/nh", "/v")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return false, "", ""
	}
	d := charmap.CodePage850.NewDecoder()
	outDecoded, err := d.Bytes(out)
	r := csv.NewReader(strings.NewReader(string(outDecoded)))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
		return false, "", ""
	}
	titlePos := 8
	processPos := 0
	for _, record := range records {
		if record[processPos] == "Spotify.exe" && strings.Contains(record[titlePos], "-") {
			parts := strings.Split(record[titlePos], "-")
			fmt.Printf("Artist: %s - Song : %s\n", parts[0], parts[1])
			return true, parts[0], parts[1]
		}
	}
	return false, "", ""
}

func (h *Spotify) GetMetadata(newSong chan *Song) {
	metadata := &Song{}
	found, title, artist := getDataFromTasklist()
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
			ok, _, _ := getDataFromTasklist()
			if ok {
				changes <- "yes"
			}
		}
	}
}

func GetDbPath() string {
	return filepath.join(os.Getenv("APPDATA"), "lyricfier")
}
