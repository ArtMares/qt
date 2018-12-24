package controller

import (
	"github.com/ArtMares/qt/core"
	"github.com/ArtMares/qt/gui"
)

type logoController struct {
	core.QObject

	_ func() `signal:"clicked,auto"`
}

func (c *logoController) clicked() {
	gui.QDesktopServices_OpenUrl(core.NewQUrl3("https://sia.tech", 0))
}
