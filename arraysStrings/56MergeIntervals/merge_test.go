package mergeintervals_test

import (
	"reflect"
	"testing"

	mergeintervals "github.com/sunilgopinath/algomap/arraysStrings/56MergeIntervals"
)

func TestMer(t *testing.T) {
	type test struct {
		input [][]int
		want  [][]int
	}

	tests := []test{
		{
			input: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			input: [][]int{{1, 4}, {4, 5}},
			want:  [][]int{{1, 5}},
		},
	}

	for _, tc := range tests {
		got := mergeintervals.Merge(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Merge(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}