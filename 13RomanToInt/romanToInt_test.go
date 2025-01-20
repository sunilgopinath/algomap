package romantoint_test

import (
	"testing"

	romantoint "github.com/sunilgopinath/algomap/13RomanToInt"
)

func TestRomanToInt(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{
			input: "III",
			want:  3,
		},
		{
			input: "IV",
			want:  4,
		},
		{
			input: "IX",
			want:  9,
		},
		{
			input: "LVIII",
			want:  58,
		},
		{
			input: "MCMXCIV",
			want:  1994,
		},
	}

	for _, tc := range tests {
		got := romantoint.RomanToInt(tc.input)
		if got != tc.want {
			t.Errorf("RomanToInt(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}