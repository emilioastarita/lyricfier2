package main

import (
	"flag"
	"github.com/emilioastarita/lyricfier2/internal/lyricfier"
	"github.com/pkg/browser"
)

var lyricfierMain *lyricfier.Main

//go:generate esc -o ../internal/lyricfier/static.go -pkg lyricfier static/
func main() {

	address := flag.String("address", "localhost", "Bind address")
	port := flag.String("port", "2387", "Bind port")

	flag.Parse()

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

	go func() {
		browser.OpenURL("http://" + *address + ":" + *port)
	}()

	lyricfierMain.StartServer(*address + ":" + *port)
}
