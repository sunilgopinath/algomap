package validsudoku_test

import (
	"testing"

	validsudoku "github.com/sunilgopinath/algomap/hashmaps/36ValidSudoku"
)

func TestIsValidSudoku(t *testing.T) {
	tests := []struct {
		board [][]byte
		want  bool
	}{
		{
			[][]byte{
				[]byte("53..7...."),
				[]byte("6..195..."),
				[]byte(".98....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8..79"),
			},
			true,
		},
		{
			[][]byte{
				[]byte("53..7...."),
				[]byte("6..195..."),
				[]byte(".98....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8..78"),
			},
			false,
		},
	}
	for _, tt := range tests {
		if got := validsudoku.IsValidSudoku(tt.board); got != tt.want {
			t.Errorf("IsValidSudoku(%v) = %v, want %v", tt.board, got, tt.want)
		}
	}
}