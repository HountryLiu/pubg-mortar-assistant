package util

import (
	"github.com/HountryLiu/pubg-mortar-assistant/screen"
	"github.com/kbinani/screenshot"
)

func GetScreenSize() (int, int) {
	pubg := screen.GetPubgInstance()
	bounds := screenshot.GetDisplayBounds(pubg.DisplayIndex)
	return bounds.Max.X - bounds.Min.X, bounds.Max.Y - bounds.Min.Y
}
