package controller

import (
	"github.com/ArtMares/qt/core"

	"github.com/ArtMares/qt/internal/examples/showcases/sia/controller"
)

type StatusController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"balance"`
	_ string `property:"delta"`

	_ func(interface{}) `signal:"wallet,<-(controller.Controller.WalletChanged)"`
}

func (c *StatusController) init() {
	if controller.DEMO {
		c.wallet(nil)
	}
}

func (c *StatusController) wallet(wgI interface{}) {
	if controller.DEMO {
		c.SetBalance("12345 SC")
		c.SetDelta("67890 SC")
	}
}
