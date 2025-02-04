package validpalindrome

import "unicode"

func IsValidPalindrome(s string) bool {

	lPtr, rPtr := 0, len(s)-1
	for lPtr < rPtr {
		if !isAlphaNumeric(s[lPtr]) {
			lPtr++
			continue
		} else if !isAlphaNumeric(s[rPtr]) {
			rPtr--
			continue
		}
		if unicode.ToLower(rune(s[lPtr])) != unicode.ToLower(rune(s[rPtr])) {
			return false
		}
		lPtr++
		rPtr--
	}
	return true
}

func isAlphaNumeric(b byte) bool {
	return unicode.IsLetter(rune(b)) || unicode.IsDigit(rune(b))
}