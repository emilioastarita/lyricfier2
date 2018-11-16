package main

import (
	"github.com/emilioastarita/lyricfier2/internal/gui"
	"github.com/emilioastarita/lyricfier2/internal/lyricfier"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
	"runtime"
	"time"
)

const (
	winWidth  = 400
	winHeight = 500
)

func init() {
	runtime.LockOSThread()
}

var lyricfierMain *lyricfier.Main

func main() {

	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	win, err := glfw.CreateWindow(winWidth, winHeight, "Lyricfier 2", nil, nil)
	if err != nil {
		panic(err)
	}
	win.MakeContextCurrent()

	width, height := win.GetSize()
	if err := gl.Init(); err != nil {
		panic(err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))

	ctx := nk.NkPlatformInit(win, nk.PlatformInstallCallbacks)

	atlas := nk.NewFontAtlas()
	nk.NkFontStashBegin(&atlas)
	sansFont := nk.NkFontAtlasAddDefault(atlas, 16, nil)
	nk.NkFontStashEnd()
	if sansFont != nil {
		nk.NkStyleSetFont(ctx, sansFont.Handle())
	}

	exitC := make(chan struct{}, 1)
	doneC := make(chan struct{}, 1)

	fpsTicker := time.NewTicker(time.Second / 30)

	lyricfierMain = &lyricfier.Main{}
	lyricfierMain.Init()
	lyricfierMain.Lookup()
	for {
		select {
		case <-exitC:
			nk.NkPlatformShutdown()
			glfw.Terminate()
			fpsTicker.Stop()
			close(doneC)
			return
		case <-fpsTicker.C:
			if win.ShouldClose() {
				close(exitC)
				continue
			}
			glfw.PollEvents()
			gui.GfxMain(win, ctx, lyricfierMain)
		case <-lyricfierMain.Conn.Changes:
			lyricfierMain.Lookup()
		case s := <-lyricfierMain.NewSongChannel:
			lyricfierMain.ReceiveSong(s)
		case l := <-lyricfierMain.LyricSearchChannel:
			lyricfierMain.ReceiveLyric(l)
		}

	}
}
