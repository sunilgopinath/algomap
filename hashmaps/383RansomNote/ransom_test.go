package ransomnote_test

import (
	"testing"

	ransomnote "github.com/sunilgopinath/algomap/hashmaps/383RansomNote"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}
	for _, tt := range tests {
		if got := ransomnote.CanConstruct(tt.ransomNote, tt.magazine); got != tt.want {
			t.Errorf("CanConstruct(%v, %v) = %v, want %v", tt.ransomNote, tt.magazine, got, tt.want)
		}
	}
}