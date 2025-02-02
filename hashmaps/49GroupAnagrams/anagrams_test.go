package anagrams_test

import (
	"reflect"
	"sort"
	"testing"

	anagrams "github.com/sunilgopinath/algomap/hashmaps/49GroupAnagrams"
)

// Sorts an array of string slices
func sortGroups(groups [][]string) {
	// Sort individual anagram groups
	for i := range groups {
		sort.Strings(groups[i])
	}
	// Sort the outer slice
	sort.Slice(groups, func(i, j int) bool {
		return groups[i][0] < groups[j][0]
	})
}

func TestGroupAnagrams(t *testing.T) {
	testTable := []struct {
		input  []string
		output [][]string
	}{
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

	for _, tt := range testTable {
		got := anagrams.GroupAnagrams(tt.input)

		// Sort both expected and actual results
		sortGroups(got)
		sortGroups(tt.output)

		if !reflect.DeepEqual(got, tt.output) {
			t.Errorf("GroupAnagrams(%v) = %v, want %v", tt.input, got, tt.output)
		}
	}
}
