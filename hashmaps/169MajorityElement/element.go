package majorityelement

func MajorityElement(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

func MajorityElementBoyersMoore(nums []int) int {
	ans := 0
	count := 0
	for _, v := range nums {
		// is our count 0?
		if count == 0 {
			ans = v
		}
		// if the current element is the same as the majority element
		if v == ans {
			count++
		} else {
			count--
		}
	}
	return ans
}