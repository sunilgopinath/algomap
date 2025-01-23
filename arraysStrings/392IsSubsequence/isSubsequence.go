package subsequence

// https://leetcode.com/problems/is-subsequence/description/

func IsSubsequence(s string, t string) bool {
	sPtr, tPtr := 0, 0
	for sPtr < len(s) && tPtr < len(t) {
		if s[sPtr] == t[tPtr] {
			sPtr++
		}
		tPtr++
	}
	return sPtr == len(s)
}