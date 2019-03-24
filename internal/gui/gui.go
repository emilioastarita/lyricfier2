package gui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	. "github.com/therecipe/qt/qml"
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
	ctx.SetRunning(false)
}

var guiSong = NewCtxObject(nil)

func Main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.SetWindowIcon(gui.NewQIcon5(":/qml/icon.png"))
	engine := NewQQmlApplicationEngine(nil)
	context := engine.RootContext()
	context.SetContextProperty("song", guiSong)
	engine.Load(core.NewQUrl3("qrc:/qml/application.qml", 0))
	gui.QGuiApplication_Exec()
}
