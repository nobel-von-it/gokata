package files

import (
	"fmt"
	"strings"
)

func MultiTable(n int) string {
	arr := make([]string, 10)
	for i := 1; i <= 10; i++ {
		arr[i-1] = fmt.Sprintf("%d * %d = %d", i, n, n*i)
	}
	return strings.Join(arr, "\n")
}
