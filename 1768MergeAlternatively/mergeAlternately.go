package mergealternately

// https://leetcode.com/problems/merge-strings-alternately/description/

func MergeAlternately(word1 string, word2 string) string {
	var result string
	var i, j int
	for i < len(word1) && j < len(word2) {
		result += string(word1[i])
		result += string(word2[j])
		i++
		j++
	}
	for i < len(word1) {
		result += string(word1[i])
		i++
	}
	for j < len(word2) {
		result += string(word2[j])
		j++
	}
	return result
}