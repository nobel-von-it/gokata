package files

import "fmt"

func countSheep(n int) (str string) {
	for i := range make([]int, n) {
		str += fmt.Sprintf("%d sheep...", i+1)
	}
	return
}
