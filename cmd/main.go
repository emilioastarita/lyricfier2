package main

import (
	"github.com/emilioastarita/lyricfier2/internal/lyricfier"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

var lyricfierMain *lyricfier.Main

func main() {
	lyricfierMain = &lyricfier.Main{}
	lyricfierMain.Init()
	lyricfierMain.Lookup()
	go func() {
		for {
			select {
			case <-lyricfierMain.Detector.Changes:
				lyricfierMain.Lookup()
			case s := <-lyricfierMain.NewSongChannel:
				lyricfierMain.ReceiveSong(s)
			case l := <-lyricfierMain.LyricSearchChannel:
				lyricfierMain.ReceiveLyric(l)
			}
		}
	}()
}
