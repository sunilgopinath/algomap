package longestsubsequence

func LongestConsecutive(nums []int) int {
    m := make(map[int]bool)
    for _, v := range nums {
        m[v] = true
    }

    maxSeq := 0
    for _, v := range nums {
        if m[v-1] {
            // not start of sequence, skip
            continue
        }
        seq, tmp := 1, v+1
        for m[tmp] {
            seq++
            tmp++
        }
        maxSeq = max(maxSeq, seq)
        if seq > len(nums) / 2 {
            break
        }
    }
    return maxSeq
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}