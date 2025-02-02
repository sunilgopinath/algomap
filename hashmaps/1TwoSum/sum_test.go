package sum_test

import (
	"reflect"
	"testing"

	sum "github.com/sunilgopinath/algomap/hashmaps/1TwoSum"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		if got := sum.TwoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("TwoSum(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}