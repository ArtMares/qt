package view

import (
	"github.com/ArtMares/qt/quick"

	_ "github.com/ArtMares/qt/internal/examples/showcases/sia/view/controller"
)

func init() { viewTemplate_QmlRegisterType2("ViewTemplate", 1, 0, "ViewTemplate") }

type viewTemplate struct {
	quick.QQuickItem

	_ func(b bool) `signal:"Blur,<-(controller.NewViewController(nil))"`
}
