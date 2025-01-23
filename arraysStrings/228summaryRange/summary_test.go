package summaryrange_test

import (
	"reflect"
	"testing"

	summaryrange "github.com/sunilgopinath/algomap/arraysStrings/228summaryRange"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		input []int
		want  []string
	}{
		{
			input: []int{0, 1, 2, 4, 5, 7},
			want:  []string{"0->2", "4->5", "7"},
		},
		{
			input: []int{0, 2, 3, 4, 6, 8, 9},
			want:  []string{"0", "2->4", "6", "8->9"},
		},
	}
	for _, d := range tests {
		got := summaryrange.SummaryRanges(d.input)
		if !reflect.DeepEqual(got, d.want) {
			t.Errorf("SummaryRanges(%v) = %v; want %v", d.input, got, d.want)
		}
	}
}