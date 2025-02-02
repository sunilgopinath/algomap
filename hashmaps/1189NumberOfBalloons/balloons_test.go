package numberofballoons_test

import (
	"testing"

	numberofballoons "github.com/sunilgopinath/algomap/hashmaps/1189NumberOfBalloons"
)

func TestMaxNumberOfBalloons(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{"nlaebolko", 1},
		{"loonbalxballpoon", 2},
		{"leetcode", 0},
	}
	for _, tt := range tests {
		if got := numberofballoons.MaxNumberOfBalloons(tt.text); got != tt.want {
			t.Errorf("MaxNumberOfBalloons(%v) = %v, want %v", tt.text, got, tt.want)
		}
	}
}