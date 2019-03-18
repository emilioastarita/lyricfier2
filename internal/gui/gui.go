package gui

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
	"os"
)

func SetLyric(text string) {
	guiSong.SetLyric(text)
}

func SetArtist(text string) {
	guiSong.SetArtist(text)
}

func SetTitle(text string) {
	guiSong.SetTitle(text)
}

func SetRunning(running bool) {
	fmt.Printf("Setting running to %v\n", running)
	guiSong.SetRunning(running)
}

type CtxObject struct {
	core.QObject
	_ string `property:"title"`
	_ string `property:"artist"`
	_ string `property:"lyric"`
	_ bool   `property:"running"`
	_ func() `constructor:"init"`
}

func (ctx *CtxObject) init() {
	ctx.ConnectRunningChanged(func(boolProp bool) {
		fmt.Println(" go: changed bool ->", boolProp)
	})
	ctx.SetRunning(false)
}

var guiSong = NewCtxObject(nil)

func Main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	app := widgets.NewQApplication(len(os.Args), os.Args)
	view := quick.NewQQuickView(nil)
	view.SetTitle("Lyricfier 2")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.RootContext().SetContextProperty("song", guiSong)
	view.SetSource(core.NewQUrl3("qrc:/qml/application.qml", 0))
	view.Show()
	app.Exec()
}
