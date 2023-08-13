package files

import (
	"regexp"
	"strings"
)

func ReplaceDots(s string) string {
	return strings.ReplaceAll(s, ".", "-")
}

// or (with regexp). I don't know

func RepDots(s string) string {
	return regexp.MustCompile(`\.`).ReplaceAllString(s, "-")
}
