package controller

import (
	"github.com/ArtMares/qt/core"

	_ "github.com/ArtMares/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

var ButtonController *buttonController

type buttonController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"clicked,->(controller.Controller.Show)"`
}

func (c *buttonController) init() {
	ButtonController = c
}
