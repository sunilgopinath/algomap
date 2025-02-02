package numberofballoons

func MaxNumberOfBalloons(text string) int {
	freq := [26]int{}

	for _, v := range text {
		freq[v-'a']++
	}
    return min(freq['b'-'a'], freq['a'-'a'], freq['l'-'a']/2, freq['o'-'a']/2, freq['n'-'a'])
}

func min(aa ...int) int {
	minVal := aa[0]
	for _, a := range aa {
		if a < minVal {
			minVal = a
		}
	}
	return minVal
}