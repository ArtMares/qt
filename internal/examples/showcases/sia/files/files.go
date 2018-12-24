package files

import (
	"github.com/ArtMares/qt/core"
	"github.com/ArtMares/qt/quick"

	"github.com/ArtMares/qt/internal/examples/showcases/sia/files/controller"
	"github.com/ArtMares/qt/internal/examples/showcases/sia/files/dialog"
)

func init() { filesTemplate_QmlRegisterType2("FilesTemplate", 1, 0, "FilesTemplate") }

type filesTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"FilesModel"`
}

func (t *filesTemplate) init() {
	c := controller.NewFilesController(nil)

	t.SetFilesModel(c.Model().Filter)

	//needed here, because those are non qml views
	dialog.NewFilesUploadTemplate(nil)
	dialog.NewFolderUploadTemplate(nil)
}
