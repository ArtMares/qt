package top

import (
	"github.com/ArtMares/qt/quick"

	_ "github.com/ArtMares/qt/internal/examples/showcases/sia/view/top/controller"
)

func init() { searchTemplate_QmlRegisterType2("TopTemplate", 1, 0, "SearchTemplate") }

type searchTemplate struct {
	quick.QQuickItem

	_ func(string) `signal:"search,->(controller.NewSearchController(nil))"`
}
