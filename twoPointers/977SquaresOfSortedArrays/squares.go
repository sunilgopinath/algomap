package squaresofsortedarrays

// leetcode #977
// SquaresofsortedarraysortedSquares returns the squares of the sorted array
func SquaresofsortedarraysortedSquares(nums []int) []int {

	lPtr, rPtr := 0, len(nums)-1
	result := make([]int, len(nums))
	for i := len(nums)-1; i >= 0; i-- {
		if abs(nums[lPtr]) > abs(nums[rPtr]) {
			result[i] = nums[lPtr] * nums[lPtr]
			lPtr++
		} else {
			result[i] = nums[rPtr] * nums[rPtr]
			rPtr--
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
