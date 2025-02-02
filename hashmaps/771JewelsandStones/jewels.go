package jewelsandstones

func NumJewelsInStones(jewels string, stones string) int {
    // make a map of jewels
	// loop through stones and check if it is a jewel
	m := make(map[rune]bool)
	for _, v := range jewels {
		m[v] = true
	}

	res := 0
	for _, v := range stones {
		if _, ok := m[v]; ok {
			res++
		}
	}
	return res
}