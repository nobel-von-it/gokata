package files

func Points(games []string) (count int) {
	for _, game := range games {
		if game[0] > game[2] {
			count += 3
		}
		if game[0] == game[2] {
			count += 1
		}
	}
	return
}

// golang is bad for fast algorithms olympiad(, maybe)
