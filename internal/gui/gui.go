package gui

import (
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

type CtxObject struct {
	core.QObject
	_ string `property:"title"`
	_ string `property:"artist"`
	_ string `property:"lyric"`
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
