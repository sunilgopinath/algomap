package summaryrange

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/summary-ranges/
func SummaryRanges(nums []int) []string {

	if len(nums) == 0 {
		return []string{}
	}

	res := make([]string, 0)
	start := nums[0]
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i] - nums[i-1] != 1 { // reached the end
			if start == nums[i-1] {
				res = append(res, strconv.Itoa(start))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			if i < len(nums) {
				start = nums[i]
			}
		}
	}

	return res
}