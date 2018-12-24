package terminal

import (
	"github.com/ArtMares/qt/quick"

	_ "github.com/ArtMares/qt/internal/examples/showcases/sia/terminal/controller"
)

func init() { terminalTemplate_QmlRegisterType2("TerminalTemplate", 1, 0, "TerminalTemplate") }

type terminalTemplate struct {
	quick.QQuickItem

	_ func(cmd string) string `slot:"command,->(controller.NewTerminalController(nil))"`
}
