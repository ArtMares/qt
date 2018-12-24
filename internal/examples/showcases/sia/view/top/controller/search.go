package controller

import (
	"github.com/ArtMares/qt/core"

	"github.com/ArtMares/qt/internal/examples/showcases/sia/files/controller"
	lcontroller "github.com/ArtMares/qt/internal/examples/showcases/sia/view/left/controller"
)

type searchController struct {
	core.QObject

	_ func(string) `signal:"search,auto"`
}

func (c *searchController) search(name string) {
	lcontroller.LeftController.Clicked("files")
	controller.FilesController.Model().Filter.SetFilterFixedString(name)
}
