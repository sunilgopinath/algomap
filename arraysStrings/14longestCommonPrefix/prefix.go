package longestcommonprefix

// https://leetcode.com/problems/longest-common-prefix/
func LongestCommonPrefix(strs []string) string {
    // initiate result to ""
    // for loop for length of first string
    // i is first char of first string
    // compare i with ith value in the rest of the strings
    // if i does not match, return result
    // if i matches, add i to result
    // continue
    result := ""
    for i := 0; i < len(strs[0]); i++ {
        for j := 1; j < len(strs); j++ {
            if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
                return result
            }
        }
        result += string(strs[0][i])
    }
    return result
}