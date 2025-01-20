package baseballgame_test

import (
	"testing"

	baseballgame "github.com/sunilgopinath/algomap/682BaseballGame"
)



func TestCalPoints(t *testing.T) {
	type test struct {
		input []string
		want   int
	}

	tests := []test{
		{
			input: []string{"5","2","C","D","+"},
			want: 30,
		},
		{
			input: []string{"5","-2","4","C","D","9","+","+"},
			want: 27,
		},
	}
	for _, d := range tests {
		got := baseballgame.CalPoints(d.input)
		if got != d.want {
			t.Errorf("CalPoi(%v) = %v; want %v", d.input, got, d.want)
		}
	}
}