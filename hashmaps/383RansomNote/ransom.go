package ransomnote

func CanConstruct(ransomNote string, magazine string) bool {
	freq := [26]int{}
	for _, v := range magazine {
		freq[v-'a']++
	}
	for _, v := range ransomNote {
		freq[v-'a']--
		if freq[v-'a'] < 0 {
			return false
		}
	}
	return true

}