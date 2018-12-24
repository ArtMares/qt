//source: http://doc.qt.io/qt-5/qtsensors-accelbubble-example.html

package main

import (
	"os"

	"github.com/ArtMares/qt/core"
	"github.com/ArtMares/qt/gui"
	"github.com/ArtMares/qt/qml"
	_ "github.com/ArtMares/qt/sensors"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("qrc:///accelbubble.qml", 0))

	gui.QGuiApplication_Exec()
}
