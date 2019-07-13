package lyricfier

import "runtime"

type DetectCurrentSong struct {
	platform string
	Conn     Spotify
	Changes  chan string
}

func (h *DetectCurrentSong) Init() {
	h.platform = runtime.GOOS
	h.Changes = make(chan string)
	h.Conn = Spotify{}
	h.Conn.Init()
	go h.Conn.Ticker(h.Changes)
}

func (h *DetectCurrentSong) GetMetadata(newSong chan *Song) {
	h.Conn.GetMetadata(newSong)
}
