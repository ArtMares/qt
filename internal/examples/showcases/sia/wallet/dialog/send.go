package dialog

import (
	"github.com/ArtMares/qt/core"

	_ "github.com/ArtMares/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

func init() { sendTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "SendTemplate") }

type sendTemplate struct {
	dialogTemplate

	_ func(string, string) *core.QVariant `slot:"send,->(controller.Controller)"`
}
