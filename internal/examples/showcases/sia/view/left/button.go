package left

import (
	"github.com/ArtMares/qt/quick"

	_ "github.com/ArtMares/qt/internal/examples/showcases/sia/view/left/controller"
)

func init() { buttonTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "ButtonTemplate") }

type buttonTemplate struct {
	quick.QQuickItem

	_ func(string) `signal:"clicked,->(controller.NewButtonController(nil))"`
}
