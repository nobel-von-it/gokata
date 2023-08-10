package files

func Feast(b, d string) bool {
	return b[0] == d[0] && b[len(b)-1] == d[len(d)-1]
}
