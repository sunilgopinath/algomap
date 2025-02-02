package anagrams

import "sort"

// leetcode #49. Group Anagrams
// https://leetcode.com/problems/group-anagrams/
func GroupAnagrams(strs []string) [][]string {

	// declare a map with key, frequency of characters in the string
	m := make(map[[26]int][]string)

	// loop through strings and create the freq map
	for _, s := range strs {
		freq := [26]int{}

		for _, c := range s {
			freq[c-'a']++
		}
		m[freq] = append(m[freq], s)
	}

	// create a result slice
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func GroupAnagramsV2(strs []string) [][]string {
    m := make(map[string][]string)

    for _, str := range strs {
        sortedStr := sortString(str) // Sort the string
        m[sortedStr] = append(m[sortedStr], str)
    }

    // Collect the results
    var res [][]string
    for _, group := range m {
        res = append(res, group)
    }

    return res
}

// Sorts the characters in a string
func sortString(s string) string {
    chars := []byte(s) // Convert string to byte slice
    sort.Slice(chars, func(i, j int) bool {
        return chars[i] < chars[j]
    })
    return string(chars) // Convert back to string
}