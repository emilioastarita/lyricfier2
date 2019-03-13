package main

import (
	"github.com/emilioastarita/lyricfier2/internal/gui"
	"github.com/emilioastarita/lyricfier2/internal/lyricfier"
	"runtime"
	"time"
)

func init() {
	runtime.LockOSThread()
}

var lyricfierMain *lyricfier.Main

func main() {

	exitC := make(chan struct{}, 1)
	doneC := make(chan struct{}, 1)

	fpsTicker := time.NewTicker(time.Second / 30)

	lyricfierMain = &lyricfier.Main{}
	lyricfierMain.Init()
	lyricfierMain.Lookup()
	go func() {
		for {
			select {
			case <-exitC:
				close(doneC)
				return
			case <-fpsTicker.C:

			case <-lyricfierMain.Detector.Changes:
				lyricfierMain.Lookup()
			case s := <-lyricfierMain.NewSongChannel:
				lyricfierMain.ReceiveSong(s)
			case l := <-lyricfierMain.LyricSearchChannel:
				lyricfierMain.ReceiveLyric(l)
			}

		}
	}()
	gui.Main()
}
