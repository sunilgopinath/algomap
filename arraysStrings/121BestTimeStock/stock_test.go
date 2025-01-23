package besttimestock_test

import (
	"testing"

	besttimestock "github.com/sunilgopinath/algomap/arraysStrings/121BestTimeStock"
)


func TestMaxProfit(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{7, 1, 5, 3, 6, 4},
			want:  5,
		},
		{
			input: []int{7, 6, 4, 3, 1},
			want:  0,
		},
		{
			input: []int{1, 2, 3, 4, 5},
			want:  4,
		},
	}
	for _, d := range tests {
		got := besttimestock.MaxProfit(d.input)
		if got != d.want {
			t.Errorf("maxProfit(%v) = %v; want %v", d.input, got, d)
		}
	}
}