package main

import (
	"os"

	"github.com/ArtMares/qt/core"
	"github.com/ArtMares/qt/gui"
	"github.com/ArtMares/qt/quick"

	_ "github.com/ArtMares/qt/internal/examples/qml/extending/components/test_dir_2/component"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/app.qml", 0))
	view.Show()

	gui.QGuiApplication_Exec()
}
