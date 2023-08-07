package util

import "github.com/kbinani/screenshot"

func GetScreenSize(displayIndex int) (int, int) {
	bounds := screenshot.GetDisplayBounds(displayIndex)
	return bounds.Max.X - bounds.Min.X, bounds.Max.Y - bounds.Min.Y
}
