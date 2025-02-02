package anagrams_test

import (
	"reflect"
	"testing"

	anagrams "github.com/sunilgopinath/algomap/hashmaps/49GroupAnagrams"
)

func TestGroupAnagrams(t *testing.T) {
	testTable := []struct{
		input []string
		output [][]string
	}{	// Define a slice of struct to hold test cases	
		{	
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, 
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			[]string{"a"}, 
			[][]string{{"a"}},
		},
		{
			[]string{"a", "b"}, 
			[][]string{{"a"}, {"b"}},
		},
	}
	for _, tt := range testTable {	// Loop through the test cases
		if got := anagrams.GroupAnagrams(tt.input); !reflect.DeepEqual(got, tt.output) {
			t.Errorf("GroupAnagrams(%v) = %v, want %v", tt.input, got, tt.output)
		}
	}
}