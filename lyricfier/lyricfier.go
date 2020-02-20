package main

import (
	"flag"
	"fmt"
	"github.com/emilioastarita/lyricfier2/icon"
	"github.com/emilioastarita/lyricfier2/internal/lyricfier"
	"github.com/getlantern/systray"
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

	onOpenBrowser := func() {
		url := "http://" + *address + ":" + *port
		browser.OpenURL(url)
		fmt.Printf("=> Open your browser at: \033[1;36m%s\033[0m\n", url)
	}

	onExit := func() {}
	onReady := func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
		systray.SetTooltip("Lyricfier")
		url := "http://" + *address + ":" + *port
		mOpenBrowser := systray.AddMenuItem("Open Lyricfier", "Or visit with your browser at:"+url)
		mQuitOrig := systray.AddMenuItem("Quit", "Quit lyricfier")
		go func() {
			for {
				select {
				case <-mOpenBrowser.ClickedCh:
					go onOpenBrowser()
				case <-mQuitOrig.ClickedCh:
					systray.Quit()
				}
			}
		}()
		go onOpenBrowser()

		go func() {
			lyricfierMain.StartServer(*address + ":" + *port)
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
	systray.RunWithAppWindow("Lyricfier", 0, 0, onReady, onExit)
}
