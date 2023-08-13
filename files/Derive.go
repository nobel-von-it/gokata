package files

import "fmt"

func Derive(c, e int) string {
	return fmt.Sprintf("%dx^%d", c*e, e-1)
}
