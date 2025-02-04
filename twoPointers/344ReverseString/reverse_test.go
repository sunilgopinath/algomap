package stringreverse_test

import (
	"testing"

	stringreverse "github.com/sunilgopinath/algomap/twoPointers/344ReverseString"
)

func TestReverseString(t *testing.T) {
	testTable := []struct {
		name string
		s    []byte
		want []byte
	}{
		{
			name: "test 1",
			s:    []byte{'h', 'e', 'l', 'l', 'o'},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "test 2",
			s:    []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			stringreverse.ReverseString(tt.s)
			if string(tt.s) != string(tt.want) {
				t.Errorf("got: %v, want: %v", tt.s, tt.want)
			}
		})
	}
}