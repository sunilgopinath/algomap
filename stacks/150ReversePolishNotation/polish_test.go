package reversepolishnotation_test

import (
	"testing"

	reversepolishnotation "github.com/sunilgopinath/algomap/stacks/150ReversePolishNotation"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{
			input: []string{"2","1","+","3","*"},
			want:  9,
		},
		{
			input: []string{"4","13","5","/","+"},
			want:  6,
		},
		{
			input: []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"},
			want:  22,
		},
	}
	for _, d := range tests {
		got := reversepolishnotation.EvalRPN(d.input)
		if got != d.want {
			t.Errorf("EvalRPN(%v) = %v; want %v", d.input, got, d.want)
		}
	}
}