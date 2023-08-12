package files

import "sort"

func ReverseSeq(n int) (res []int) {
	res = make([]int, n)
	for i := range res {
		res[i] = i + 1
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	} // 2.5 sec (best practices)
	sort.Sort(sort.Reverse(sort.IntSlice(res))) // 3.5 sec
	return
}
