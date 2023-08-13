package files

func FindMultiples(n, l int) (a []int) {
	for i := n; i <= l; i += n {
		a = append(a, i)
	}
	return
}
