package resource

// This file embeds all the resources used by the program.

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed fire.png
var embedDesktopIcon []byte
var DesktopIcon = fyne.NewStaticResource("DesktopIcon", embedDesktopIcon)
