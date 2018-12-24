//source: https://github.com/benlau/quickflux

package main

import (
	"os"

	"github.com/ArtMares/qt/core"
	"github.com/ArtMares/qt/gui"
	"github.com/ArtMares/qt/qml"

	_ "github.com/ArtMares/qt/internal/examples/3rdparty/quickflux/quickflux"
)

func main() {

	os.Setenv("QML_DISABLE_DISK_CACHE", "true")

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	app := qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("qrc:/main.qml", 0))

	gui.QGuiApplication_Exec()
}
