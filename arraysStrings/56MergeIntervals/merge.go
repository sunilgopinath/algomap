package mergeintervals

import "sort"

func Merge(intervals [][]int) [][]int {
	// sort the intervals

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// initialize result 2 d array
	var res [][]int

	// set inteval start, end, know input is at least 1
	start := intervals[0][0]
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end { // overlap, merge
			end = max(end, intervals[i][1])
		} else {
			res = append(res, []int{start, end })
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	res = append(res, []int{start, end})
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}