package containsduplicates

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}
	return false
}