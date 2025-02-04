package validpalindrome_test

import (
	"testing"

	validpalindrome "github.com/sunilgopinath/algomap/twoPointers/125ValidPalindrome"
)

func TestIsValidPalindrome(t *testing.T) {
	testTable := []struct{
		name string
		sentence string
		want bool
	}{
		{
			name: "test-1",
			sentence: "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "test-2",
			sentence: "race a car",
			want: false,
		},
	}
	for _, tt := range testTable {
		got := validpalindrome.IsValidPalindrome(tt.sentence)
		if got != tt.want {
			t.Errorf("got: %v, want: %v", got, tt.want)
		}
	}
}