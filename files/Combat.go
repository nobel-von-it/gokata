package files

import "math"

func Combat(h, d float64) (res float64) {
	if res = h - d; res > 0 {
		return
	} else {
		res = 0
	}
	return
}

// or simple

func Combatpp(h, d float64) float64 {
	return math.Max(h-d, 0)
}
