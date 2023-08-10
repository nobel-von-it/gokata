package files

import (
	"math"
	"sort"
)

func ExpMatter(a, b, c int) int {
	return max(a*(b+c), a*b*c, a+b*c, (a+b)*c)
}
func ExpressionMatter(a, b, c int) int {
	return int(math.Max(float64(a*(b+c)),
		math.Max(float64(a*b*c),
			math.Max(float64(a+b*c),
				math.Max(float64((a+b)*c), float64(a+b+c))))))
}
func ExpMatterpp(a, b, c int) int {
	arr := []int{a * (b + c), a * b * c, a + b + c, (a + b) * c}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return arr[0]
}
