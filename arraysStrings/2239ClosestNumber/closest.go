package closestnumber

/*
https://leetcode.com/problems/find-closest-number-to-zero/description/

Given an integer array nums of size n, return the number with the value closest to 0 in nums.
If there are multiple answers, return the number with the largest value.
*/

func FindClosestNumber(nums []int) int {
	closest := nums[0]
	for _, num := range nums {
		if abs(num) < abs(closest) {
			closest = num
		} else if abs(num) == abs(closest) {
			closest = max(num, closest)
		}
	}
	return closest
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}