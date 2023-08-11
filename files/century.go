package files

import "math"

func Century(y int) int {
	return int(math.Ceil(float64(y) / 100))
}
