package closestnumber_test

import (
	"testing"

	closestnumber "github.com/sunilgopinath/algomap/arraysStrings/2239ClosestNumber"
)

func TestFindClosestNumber(t *testing.T) {
    type test struct {
        input []int
        want  int
    }

    tests := []test{
        {
			input: []int{-4,-2,1,4,8},
			want: 1,
		},
		{
			input: []int{2,-1,1},
			want: 1,
		},
    }

    for _, tc := range tests {
        got := closestnumber.FindClosestNumber(tc.input)
        if got != tc.want {
			t.Errorf("FindClosestNumber(%v) = %v; want %v", tc.input, got, tc.want)
		}
    }
}