package subsequence_test

import (
	"testing"

	subsequence "github.com/sunilgopinath/algomap/392IsSubsequence"
)

func TestIsSubsequence(t *testing.T) {
	type test struct {
		s    string
		t    string
		want bool
	}

	tests := []test{
		{
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
	}

	for _, tc := range tests {
		got := subsequence.IsSubsequence(tc.s, tc.t)
		if got != tc.want {
			t.Errorf("IsSubsequence(%v, %v) = %v; want %v", tc.s, tc.t, got, tc.want)
		}
	}
}