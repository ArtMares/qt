package top

import (
	"github.com/ArtMares/qt/quick"

	"github.com/ArtMares/qt/internal/examples/showcases/sia/view/top/controller"
)

func init() { lockTemplate_QmlRegisterType2("TopTemplate", 1, 0, "LockTemplate") }

type lockTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ bool `property:"locked,<-(this.c)"`

	_ func() `signal:"change,->(this.c)"`

	c *controller.LockController
}

func (t *lockTemplate) init() {
	t.c = controller.NewLockController(nil)
}
