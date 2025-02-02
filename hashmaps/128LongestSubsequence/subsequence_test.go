package longestsubsequence_test

import (
	"testing"

	longestsubsequence "github.com/sunilgopinath/algomap/hashmaps/128LongestSubsequence"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}
	for _, tt := range tests {
		if got := longestsubsequence.LongestConsecutive(tt.nums); got != tt.want {
			t.Errorf("LongestConsecutive(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}