package jewelsandstones_test

import (
	"testing"

	jewelsandstones "github.com/sunilgopinath/algomap/hashmaps/771JewelsandStones"
)

func TestNumJewelsInStones(t *testing.T) {
	tests := []struct {		// Define a slice of struct to hold test cases
		jewels string
		stones string
		want   int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}
	for _, tt := range tests {	// Loop through the test cases
		if got := jewelsandstones.NumJewelsInStones(tt.jewels, tt.stones); got != tt.want {
			t.Errorf("NumJewelsInStones(%v, %v) = %v, want %v", tt.jewels, tt.stones, got, tt.want)
		}
	}
}