package util

import "image"

func GetScreenSize(bounds image.Rectangle) (int, int) {
	return bounds.Max.X - bounds.Min.X, bounds.Max.Y - bounds.Min.Y
}
