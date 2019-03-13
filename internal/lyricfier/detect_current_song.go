package lyricfier

import "runtime"

type DetectCurrentSong struct {
	platform  string
	ConnLinux SpotifyDbus
	Changes   chan string
}

func (h *DetectCurrentSong) Init() {
	h.platform = runtime.GOOS
	println(h.platform)
	h.Changes = make(chan string)
	if h.platform == "linux" {
		h.InitLinux()
	}
}

func (h *DetectCurrentSong) GetMetadata(newSong chan *Song) {
	if h.platform == "linux" {
		h.GetMetadataLinux(newSong)
	}
}

// linux implementation
func (h *DetectCurrentSong) InitLinux() {
	h.ConnLinux = SpotifyDbus{}
	h.ConnLinux.Init()
	go h.ConnLinux.Ticker(h.Changes)
}
func (h *DetectCurrentSong) GetMetadataLinux(newSong chan *Song) {
	h.ConnLinux.GetMetadata(newSong)
}
