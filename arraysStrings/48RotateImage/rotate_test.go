package rotateimage_test

import (
	"reflect"
	"testing"

	rotateimage "github.com/sunilgopinath/algomap/arraysStrings/48RotateImage"
)

func TestRotate(t *testing.T){
	type test struct {
		input [][]int
		want [][]int
	}
	tests := []test{
		{
			input: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			input: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			want: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}
	for _, tc := range tests {
		rotateimage.Rotate(tc.input)
		if !reflect.DeepEqual(tc.input, tc.want) {
			t.Errorf("Rotate(%v) = %v; want %v", tc.input, tc.input, tc.want)
		}
	}
}