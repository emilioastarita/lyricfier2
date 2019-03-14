package lyricfier

import (
	"github.com/emilioastarita/lyricfier2/internal/gui"
	"github.com/emilioastarita/lyricfier2/internal/search"
	"regexp"
	"strings"
)

type Song struct {
	Title      string
	Artist     string
	ArtUrl     string
	Lyric      string
	LyricFound bool
	Source     string
}

type SearchResult struct {
	Found  bool
	Lyric  string
	Source string
}

type Main struct {
	Detector           DetectCurrentSong
	NewSongChannel     chan *Song
	LyricSearchChannel chan *SearchResult
	Current            *Song
	SpotifyRunning     bool
	Searching          bool
	searchLock         bool
}

func (h *Main) Init() {
	h.Detector = DetectCurrentSong{}
	h.searchLock = false
	h.SpotifyRunning = false
	h.Detector.Init()
	h.NewSongChannel = make(chan *Song)
	h.LyricSearchChannel = make(chan *SearchResult)
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
		h.SpotifyRunning = false
		return
	}
	h.SpotifyRunning = true
	if h.Current == nil || h.Current.Title != newSong.Title {
		h.Current = newSong
		h.Current.Lyric = ""
		if h.Searching {
			return
		}
		h.Searching = true
		gui.SetArtist(newSong.Artist)
		gui.SetTitle(newSong.Title)
		gui.SetLyric("Searching...")
		go h.Search(h.LyricSearchChannel, newSong.Artist, newSong.Title)
	}
}
func (h *Main) ReceiveLyric(newLyric *SearchResult) {
	h.Searching = false
	if h.Current != nil {
		h.Current.Lyric = newLyric.Lyric
		h.Current.LyricFound = newLyric.Found
		h.Current.Source = newLyric.Source
		gui.SetLyric(newLyric.Lyric)
	}
}

func (h *Main) Search(done chan *SearchResult, artist string, title string) {
	s := &SearchResult{Found: false}
	s.Source = "Wikia"
	lyric, e := search.Wikia(artist, normalizeTitle(title))
	if e != nil || lyric != "" {
		s.Source = "Genius"
		lyric, e = search.Genius(artist, normalizeTitle(title))
	}
	if lyric != "" {
		s.Found = true
		s.Lyric = lyric
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
