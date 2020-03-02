package lyricfier

import (
	"encoding/json"
	"fmt"
	"github.com/emilioastarita/lyricfier2/internal/search"
	"os"
	"regexp"
	"strings"
)

type Song struct {
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	ArtUrl     string `json:"artUrl"`
	Lyric      string `json:"lyric"`
	LyricFound bool   `json:"found"`
	Source     string `json:"source"`
}

type SearchResult struct {
	Found  bool
	Lyric  string
	Source string
}

type AppData struct {
	Song           Song `json:"song"`
	SpotifyRunning bool
	Searching      bool
	Address        string
	RunningInSnap  bool `json:"inSnap"`
	Settings       *Settings `json:"settings"`
}

type Settings struct {
	Theme    string `json:"theme"`
	FontSize string `json: "fontSize"`
}

type Main struct {
	Detector           DetectCurrentSong
	NewSongChannel     chan *Song
	LyricSearchChannel chan *SearchResult
	AppData            *AppData
	searchLock         bool
	server             *Server
}

func (h *Main) Init() {
	h.AppData = &AppData{}
	h.ReadSettings()
	h.Detector = DetectCurrentSong{}
	h.searchLock = false
	h.AppData.SpotifyRunning = false
	h.AppData.RunningInSnap = os.Getenv("RUNNING_IN_SNAP") == "1"

	h.Detector.Init()
	h.NewSongChannel = make(chan *Song)
	h.LyricSearchChannel = make(chan *SearchResult)
	h.server = &Server{}
}
func (h *Main) StartServer(bindAddress string) {
	h.AppData.Address = bindAddress
	h.server.Init(h.AppData)
}

func (h *Main) lock() {
	h.searchLock = true
}

func (h *Main) unlock() {
	h.searchLock = false
}

func (h *Main) Lookup() {
	if h.searchLock {
		return
	}
	h.lock()
	defer h.unlock()
	go h.Detector.GetMetadata(h.NewSongChannel)
}

func (h *Main) ReceiveSong(newSong *Song) {
	if newSong == nil {
		h.AppData.SpotifyRunning = false
		return
	}

	h.AppData.SpotifyRunning = true
	if h.AppData.Song.Title != newSong.Title {
		h.AppData.Song = *newSong
		h.AppData.Song.Lyric = ""
		if h.AppData.Searching {
			return
		}
		h.AppData.Searching = true
		go h.Search(h.LyricSearchChannel, newSong.Artist, newSong.Title)
	}
	h.server.NotifyChanges()
}

func (h *Main) ReceiveLyric(newLyric *SearchResult) {
	h.AppData.Searching = false
	h.AppData.Song.Lyric = newLyric.Lyric
	h.AppData.Song.LyricFound = newLyric.Found
	h.AppData.Song.Source = newLyric.Source
	h.server.NotifyChanges()
}

func defaultSettings() *Settings {
	s := new(Settings)
	s.Theme = "default"
	s.FontSize = "1"
	return s
}

func (h *Main) ReadSettings() {
	val, err := Read(GeneralBucket, SettingsKey)
	if err != nil {
		fmt.Printf("error reading settings %v\n", err)
		h.AppData.Settings = defaultSettings()
		return
	}
	err = json.Unmarshal(val, h.AppData.Settings)
	if err != nil {
		fmt.Printf("error umarshalling saved settings %v\n", err)
		h.AppData.Settings = defaultSettings()
		return
	}
}

func (h *Main) Search(done chan *SearchResult, artist string, title string) {
	s := &SearchResult{Found: false, Source: "LocalDb"}
	key := SongKey(artist, title)
	val, e := Read(SongsBucket, key)
	var lyric = ""
	if e == nil {
		lyric = string(val)
	}
	if lyric == "" {
		s.Source = "Wikia"
		lyric, e = search.Wikia(artist, normalizeTitle(title))
		if e != nil || lyric == "" {
			s.Source = "Genius"
			lyric, e = search.Genius(artist, normalizeTitle(title))
		}
	}
	if lyric != "" {
		s.Found = true
		s.Lyric = lyric
		if s.Source != "LocalDb" {
			Write(SongsBucket, key, []byte(lyric))
		}
	}
	done <- s
}

var ignoreParts = regexp.MustCompile(`(?i)remastered|bonus track|remasterizado|live|remaster`)

func normalizeTitle(title string) string {
	parts := strings.Split(title, "-")
	if len(parts) == 2 {
		check := strings.ToLower(parts[1])
		if ignoreParts.MatchString(check) {
			return strings.Trim(parts[0], " ")
		}
	}
	return title
}
