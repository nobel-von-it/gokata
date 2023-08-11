package files

func MultipleOfIndex(ints []int) (res []int) {
	for i, el := range ints {
		if el%(i+1) == 0 {
			res = append(res, el)
		}
	}
	return
}
