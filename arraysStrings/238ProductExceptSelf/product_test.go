package productexceptself_test

import (
	"reflect"
	"testing"

	productexceptself "github.com/sunilgopinath/algomap/arraysStrings/238ProductExceptSelf"
)

func TestProductExceptSelf(t *testing.T) {
	type test struct {
		input []int
		want  []int
	}

	tests := []test{
		{
			input: []int{1, 2, 3, 4},
			want:  []int{24, 12, 8, 6},
		},
		{
			input: []int{-1, 1, 0, -3, 3},
			want:  []int{0, 0, 9, 0, 0},
		},
	}

	for _, tc := range tests {
		got := productexceptself.ProductExceptself(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ProductExceptSelf(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}