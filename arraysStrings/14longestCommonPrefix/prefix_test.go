package longestcommonprefix_test

import (
	"testing"

	longestcommonprefix "github.com/sunilgopinath/algomap/arraysStrings/14longestCommonPrefix"
)

func TestLongestCommonPrefix(t *testing.T) {
	type test struct {
		input []string
		want  string
	}

	tests := []test{
		{
			input: []string{"flower", "flow", "flight"},
			want:  "fl",
		},
		{
			input: []string{"dog", "racecar", "car"},
			want:  "",
		},
	}

	for _, tc := range tests {
		got := longestcommonprefix.LongestCommonPrefix(tc.input)
		if got != tc.want {
			t.Errorf("LongestCommonPrefix(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}