package controller

import (
	"github.com/ArtMares/qt/core"

	"github.com/ArtMares/qt/internal/examples/showcases/sia/controller"
	dcontroller "github.com/ArtMares/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

func init() {
	if false {
		controller.Controller = nil
	}
}

type LockController struct {
	core.QObject

	_ bool `property:"locked,<-(controller.Controller)"`

	_ func() `signal:"change,auto"`
}

func (c *LockController) change() {
	if c.IsLocked() {
		dcontroller.Controller.Show("unlock")
	} else {
	}
}
