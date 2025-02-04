# algomap
Doing algomap.io's 100 leetcode questions

## Answered

1. [2239. Find Closest Number to Zero](./arraysStrings/2239ClosestNumber/)
2. [1768. Merge Strings Alternatively](./arraysStrings/1768MergeAlternatively/)
3. [13. Roman To Integer](./arraysStrings/13RomanToInt/)
4. [392. Is Subsequence](./arraysStrings/392IsSubsequence/)
5. [682. Baseball Game](./stacks/682BaseballGame/)
6. [20. Valid Parentheses](./stacks/20ValidParentheses/)
7. [150. Reverse Polish Notation](./stacks/150ReversePolishNotation/)
8. [121. Best Time To Buy and Sell Stock](./stocks/121BestTimeToBuyAndSellStock/)
9. [14. Longest Common Prefix](./strings/14LongestCommonPrefix/)
10. [228. Summary Ranges](./ranges/228SummaryRanges/)
11. [238. Product Except Self](./arraysStrings/238ProductExceptSelf/)
12. [56. Merge Intervals](./arraysStrings/56MergeIntervals/)
13. [83. Remove Duplicates From Sorted List](./linkedLists/83RemoveDuplicates/)
14. [54. Spiral Matrix](./arraysStrings/54SpiralMatrix/)
15. [48. Rotate Image](./arraysStrings/48RotateImage/)
16. [771. Jewels and Stones](./hashmaps/771JewelsandStones/)
17. [217. Contains Duplicates](./hashmaps/217ContainsDuplicates/)
18. [383. Ransom Note](./hashmaps/383RansomNote/)
19. [242. Valid Anagram](./hashmaps/242ValidAnagram/)
20. [1189. Maximum Number of Balloons](./hashmaps/1189NumberOfBalloons/)
21. [1. Two Sum](./hashmaps/1TwoSum/)
22. [36. Valid Sudoku](./hashmaps/36ValidSudoku/)
23. [49. Group Anagrams](./hashmaps/49GroupAnagrams/)
24. [169. Majority Element](./hashmaps/169MajorityElement/)
25. [128. Longest Subsequence](./hashmaps/128LongestSubsequence/)
26. [977. Squares of Sorted Arrays](./twoPointers/977SquaresOfSortedArrays/)

## Run Tests
```sh
$> go test ./...
?   	github.com/sunilgopinath/algomap	[no test files]
ok  	github.com/sunilgopinath/algomap/arraysStrings/121BestTimeStock	0.523s
ok  	github.com/sunilgopinath/algomap/arraysStrings/13RomanToInt	0.732s
ok  	github.com/sunilgopinath/algomap/arraysStrings/14longestCommonPrefix	1.265s
ok  	github.com/sunilgopinath/algomap/arraysStrings/1768MergeAlternatively	1.667s
ok  	github.com/sunilgopinath/algomap/arraysStrings/2239ClosestNumber	1.478s
ok  	github.com/sunilgopinath/algomap/arraysStrings/228summaryRange	0.909s
ok  	github.com/sunilgopinath/algomap/arraysStrings/238ProductExceptSelf	1.099s
ok  	github.com/sunilgopinath/algomap/arraysStrings/392IsSubsequence	1.804s
ok  	github.com/sunilgopinath/algomap/arraysStrings/48RotateImage	1.297s
ok  	github.com/sunilgopinath/algomap/arraysStrings/54SpiralMatrix	1.303s
ok  	github.com/sunilgopinath/algomap/arraysStrings/56MergeIntervals	1.313s
ok  	github.com/sunilgopinath/algomap/hashmaps/1189NumberOfBalloons	1.307s
ok  	github.com/sunilgopinath/algomap/hashmaps/128LongestSubsequence	1.279s
ok  	github.com/sunilgopinath/algomap/hashmaps/169MajorityElement	1.262s
ok  	github.com/sunilgopinath/algomap/hashmaps/1TwoSum	1.297s
ok  	github.com/sunilgopinath/algomap/hashmaps/217ContainsDuplicates	1.197s
ok  	github.com/sunilgopinath/algomap/hashmaps/242ValidAnagram	1.209s
ok  	github.com/sunilgopinath/algomap/hashmaps/36ValidSudoku	1.161s
?   	github.com/sunilgopinath/algomap/stacks/739DailyTemperatures	[no test files]
ok  	github.com/sunilgopinath/algomap/hashmaps/383RansomNote	1.207s
ok  	github.com/sunilgopinath/algomap/hashmaps/49GroupAnagrams	1.110s
ok  	github.com/sunilgopinath/algomap/hashmaps/771JewelsandStones	1.138s
ok  	github.com/sunilgopinath/algomap/linkedLists/83RemoveDuplicates	1.142s
ok  	github.com/sunilgopinath/algomap/stacks/150ReversePolishNotation	1.107s
ok  	github.com/sunilgopinath/algomap/stacks/20Parenthesis	1.123s
ok  	github.com/sunilgopinath/algomap/stacks/682BaseballGame	1.123s
ok  	github.com/sunilgopinath/algomap/twoPointers/977SquaresOfSortedArrays	1.159s
```

### Majority Element (Boyer Moore)

#### Example Walkthrough

    - Input: [2, 2, 1, 1, 1, 2, 2]
    - Pick first number: 2 (candidate), count = 1
    - Next is 2 again: count = 2
    - Next is 1 (different number): count = 1
    - Next is 1: count = 0 (reset candidate)
    - Pick new candidate: 1, count = 1
    - Next is 2 (different number): count = 0 (reset candidate)
    - Pick new candidate: 2, count = 1
    - Next is 2: count = 2

✔ Final Majority Element: 2