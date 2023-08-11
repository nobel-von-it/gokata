package files

func CalculateYears(y int) (res [3]int) {
	switch y {
	case 1:
		res = [3]int{1, 15, 15}
	case 2:
		res = [3]int{2, 24, 24}
	default:
		res = [3]int{y, 24 + (y-2)*4, 24 + (y-2)*5}
	}
	return
}
