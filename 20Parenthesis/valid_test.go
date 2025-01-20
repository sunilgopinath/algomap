package parenthesis_test

import (
	"testing"

	parenthesis "github.com/sunilgopinath/algomap/20Parenthesis"
)


func TestIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "()",
			want:  true,
		},
		{
			input: "()[]{}",
			want:  true,
		},
		{
			input: "(]",
			want:  false,
		},
		{
			input: "([)]",
			want:  false,
		},
		{
			input: "{[]}",
			want:  true,
		},
	}
	for _, d := range tests {
		got := parenthesis.IsValid(d.input)
		if got != d.want {
			t.Errorf("IsValid(%v) = %v; want %v", d.input, got, d)
		}
	}
}