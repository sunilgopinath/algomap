package spiralmatrix_test

import (
	"reflect"
	"testing"

	spiralmatrix "github.com/sunilgopinath/algomap/arraysStrings/54SpiralMatrix"
)

func TestSpiralOrder(t *testing.T) {
	type test struct {
		input [][]int
		want  []int
	}
	tests := []test{
		{
			input: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:  []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			input: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:  []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tc := range tests {
		got := spiralmatrix.SpiralOrder(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("SpiralOrder(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}