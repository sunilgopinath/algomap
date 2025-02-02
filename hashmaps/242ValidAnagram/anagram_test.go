package validanagram_test

import (
	"testing"

	validanagram "github.com/sunilgopinath/algomap/hashmaps/242ValidAnagram"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"a", "ab", false},
	}
	for _, tt := range tests {
		if got := validanagram.IsAnagram(tt.s, tt.t); got != tt.want {
			t.Errorf("IsAnagram(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}