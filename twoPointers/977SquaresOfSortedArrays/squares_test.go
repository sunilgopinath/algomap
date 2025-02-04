package squaresofsortedarrays_test

import (
	"reflect"
	"testing"

	squaresofsortedarrays "github.com/sunilgopinath/algomap/twoPointers/977SquaresOfSortedArrays"
)

func TestSortedSquares(t *testing.T) {
	testTable := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test 1",
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "test 2",
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			got := squaresofsortedarrays.SquaresofsortedarraysortedSquares(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}