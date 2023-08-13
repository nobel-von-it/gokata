package files

import (
	"sort"
)

func MergeArrays(arr1, arr2 []int) (res []int) {
	res = append(arr1, arr2...)
	sort.Ints(res)
	for i := len(res) - 1; i > 0; i-- {
		if res[i] == res[i-1] {
			res = append(res[:i], res[i+1:]...)
		}
	}
	return
}
