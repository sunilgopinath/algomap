package sum_test

import (
	"reflect"
	"testing"

	sum "github.com/sunilgopinath/algomap/twoPointers/167TwoSumII"
)

func TestTwoSum(t *testing.T) {
	testTable := []struct {
		name string
		nums    []int
		target int
		want []int
	}{
		{
			name: "test 1",
			nums: []int{2,7,11,15},
			target: 9,
			want: []int{1, 2},
		},
		{
			name: "test 2",
			nums: []int{2,3,4},
			target: 6,
			want: []int{1, 3},
		},
	}
	for _, tt := range testTable {
		got := sum.TwoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got: %v, want: %v", got, tt.want)
		}
	}
}