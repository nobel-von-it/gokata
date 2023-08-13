package files

func Invert(arr []int) (a []int) {
	for _, el := range arr {
		a = append(a, -el)
	}
	return
}
