package validanagram

func IsAnagram(s string, t string) bool {
	// check lengths of the input strings
	if len(s) != len(t) {
		return false
	}

	// create freq map of characters
	freq := [26]int{}

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}