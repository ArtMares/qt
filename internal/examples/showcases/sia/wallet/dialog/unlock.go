package dialog

import (
	"github.com/ArtMares/qt/core"

	_ "github.com/ArtMares/qt/internal/examples/showcases/sia/wallet/dialog/controller"
)

func init() { unlockTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "UnlockTemplate") }

type unlockTemplate struct {
	dialogTemplate

	_ func(string) *core.QVariant `slot:"unlock,->(controller.Controller)"`
}
