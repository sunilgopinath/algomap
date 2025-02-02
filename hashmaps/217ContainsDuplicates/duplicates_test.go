package containsduplicates_test

import (
	"testing"

	containsduplicates "github.com/sunilgopinath/algomap/hashmaps/217ContainsDuplicates"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, tt := range tests {
		if got := containsduplicates.ContainsDuplicate(tt.nums); got != tt.want {
			t.Errorf("ContainsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}