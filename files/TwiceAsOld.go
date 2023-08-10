package files

import "math"

func TwiceAsOld(d, s int) int {
	return int(math.Abs(float64(d - (s * 2))))
}
