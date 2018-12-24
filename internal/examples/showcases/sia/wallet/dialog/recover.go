package dialog

import (
	"github.com/ArtMares/qt/core"

	_ "github.com/ArtMares/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

func init() { recoverTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "RecoverTemplate") }

type recoverTemplate struct {
	dialogTemplate

	_ func(string) *core.QVariant `slot:"recover,->(controller.Controller)"`
}
