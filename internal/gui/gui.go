package gui

import (
	"bytes"
	"fmt"
	"github.com/emilioastarita/lyricfier2/internal/lyricfier"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
	"image"
     _ "image/png"
	"strings"
)

const (
	maxVertexBuffer  = 512 * 1024
	maxElementBuffer = 128 * 1024
	winWidth         = 400
	winHeight        = 500
)

var Colors = struct {
	Primary    nk.Color
	Secondary  nk.Color
	System     nk.Color
	Body       nk.Color
	Background nk.Color
}{
	Primary:    nk.NkRgb(220, 173, 22),
	Secondary:  nk.NkRgb(6, 97, 103),
	System:     nk.NkRgb(219, 212, 193),
	Body:       nk.NkRgb(255, 255, 255),
	Background: nk.NkRgba(28, 48, 62, 255),
}

func AddLabel(ctx *nk.Context, text string, color nk.Color) {
	nk.NkLayoutRowDynamic(ctx, 20, 1)
	nk.NkLabelColored(ctx, text, nk.TextLeft, color)
}

func AddLyric(ctx *nk.Context, text string) {
	for _, line := range strings.Split(strings.TrimSuffix(text, "\n"), "\n") {
		nk.NkLabelColored(ctx, line, nk.TextLeft, Colors.Body)
	}
}

func GfxMain(win *glfw.Window, ctx *nk.Context, lyricfierMain *lyricfier.Main) {
	nk.NkPlatformNewFrame()
	w, h := win.GetSize()
	bounds := nk.NkRect(0, 0, float32(w), float32(h))
	update := nk.NkBegin(ctx, "Lyricfier", bounds, 0)

	if update > 0 {

		if lyricfierMain.SpotifyRunning == false {
			AddLabel(ctx, "Is spotify running?", Colors.System)
		}

		if lyricfierMain.Current != nil {
			win.SetTitle("Lyricfier 2 - " + lyricfierMain.Current.Title + " - " + lyricfierMain.Current.Artist)
			AddLabel(ctx, lyricfierMain.Current.Title, Colors.Primary)
			AddLabel(ctx, lyricfierMain.Current.Artist, Colors.Secondary)
			if lyricfierMain.Searching == true {
				AddLabel(ctx, "Searching...", Colors.System)
			}
			if lyricfierMain.Searching == false && lyricfierMain.Current.LyricFound == false {
				AddLabel(ctx, "Not found", Colors.Secondary)
			}
			AddLyric(ctx, lyricfierMain.Current.Lyric)
		}
	}
	nk.NkEnd(ctx)
	bg := make([]float32, 4)
	nk.NkColorFv(bg, Colors.Background)
	gl.Viewport(0, 0, int32(w), int32(h))
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(bg[0], bg[1], bg[2], bg[3])
	nk.NkPlatformRender(nk.AntiAliasingOn, maxVertexBuffer, maxElementBuffer)
	win.SwapBuffers()
}

func setIcons(win *glfw.Window, iconsBytes [][]byte) {
	images := make([]image.Image, len(iconsBytes), len(iconsBytes))
	for i := range iconsBytes {
		var (
			t string
			err error
		);
		images[i], t, err = image.Decode(bytes.NewReader(iconsBytes[i]))
		if err != nil {
			panic(err)
		}
		fmt.Println("type: ", t)
	}
	win.SetIcon(images)
}

func CreateWindow(icons [][]byte) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	win, err := glfw.CreateWindow(winWidth, winHeight, "Lyricfier 2", nil, nil)
	setIcons(win, icons)
	if err != nil {
		panic(err)
	}
	win.MakeContextCurrent()

	width, height := win.GetSize()
	if err := gl.Init(); err != nil {
		panic(err)
	}
	gl.Viewport(0, 0, int32(width), int32(height))
	return win
}
