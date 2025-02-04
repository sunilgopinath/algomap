package sum

func TwoSum(numbers []int, target int) []int {
	lPtr, rPtr := 0, len(numbers)-1

	for lPtr < rPtr {
		sum := numbers[lPtr] + numbers[rPtr]
		if sum == target {
			return []int{lPtr+1, rPtr+1} // Return 1-based indices
		} else if sum < target {
			lPtr++ // 
		} else {
			rPtr-- //
		}
	}
	return []int{-1,-1}
}
