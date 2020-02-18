package main

import (
	"flag"
	"fmt"
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
		url := "http://" + *address + ":" + *port
		browser.OpenURL(url)
		fmt.Printf("=> Open your browser at \033[1;36m%s\033[0m\n", url)
	}()

	lyricfierMain.StartServer(*address + ":" + *port)
}
