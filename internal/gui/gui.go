package gui

import (
	"github.com/emilioastarita/lyricfier2/internal/lyricfier"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
	"strings"
)

const (
	maxVertexBuffer  = 512 * 1024
	maxElementBuffer = 128 * 1024
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
