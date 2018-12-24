package top

import (
	"github.com/ArtMares/qt/quick"

	_ "github.com/ArtMares/qt/internal/examples/showcases/sia/view/top/controller"
)

func init() { colorTemplate_QmlRegisterType2("TopTemplate", 1, 0, "ColorTemplate") }

type colorTemplate struct {
	quick.QQuickItem

	_ func() `signal:"change,->(controller.NewColorController(nil))"`
}
