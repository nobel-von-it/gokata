package files

import "strings"

func Contamination(t, c string) string {
	return strings.Repeat(c, len(t))
}
