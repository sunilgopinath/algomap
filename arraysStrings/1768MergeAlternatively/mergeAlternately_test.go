package mergealternately_test

import (
	"testing"

	mergealternately "github.com/sunilgopinath/algomap/arraysStrings/1768MergeAlternatively"
)

func TestMergeAlternately(t *testing.T) {
	type test struct {
		input1 string
		input2 string
		want   string
	}

	tests := []test{
		{
			input1: "abc",
			input2: "pqr",
			want: "apbqcr",
		},
		{
			input1: "ab",
			input2: "pqrs",
			want: "apbqrs",
		},
	}

	for _, tc := range tests {
		got := mergealternately.MergeAlternately(tc.input1, tc.input2)
		if got != tc.want {
			t.Errorf("MergeAlternately(%v, %v) = %v; want %v", tc.input1, tc.input2, got, tc.want)
		}
	}
}