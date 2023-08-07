package util

import "math"

// 四舍五入取整
func Round(x float64) int {
	return int(math.Floor(x + 0.5))
}
