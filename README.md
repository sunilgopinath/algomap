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
27. [344. Reverse String](./twoPointers/344ReverseString/)
28. [167. Two Sum II](./twoPointers/167TwoSumII/)

## Run Tests
```sh
?   	github.com/sunilgopinath/algomap	[no test files]
ok  	github.com/sunilgopinath/algomap/arraysStrings/121BestTimeStock	(cached)
ok  	github.com/sunilgopinath/algomap/arraysStrings/13RomanToInt	(cached)
ok  	github.com/sunilgopinath/algomap/arraysStrings/14longestCommonPrefix	(cached)
ok  	github.com/sunilgopinath/algomap/arraysStrings/1768MergeAlternatively	(cached)
ok  	github.com/sunilgopinath/algomap/arraysStrings/2239ClosestNumber	(cached)
ok  	github.com/sunilgopinath/algomap/arraysStrings/228summaryRange	(cached)
ok  	github.com/sunilgopinath/algomap/arraysStrings/238ProductExceptSelf	(cached)
ok  	github.com/sunilgopinath/algomap/arraysStrings/392IsSubsequence	(cached)
ok  	github.com/sunilgopinath/algomap/arraysStrings/48RotateImage	(cached)
ok  	github.com/sunilgopinath/algomap/arraysStrings/54SpiralMatrix	(cached)
ok  	github.com/sunilgopinath/algomap/arraysStrings/56MergeIntervals	(cached)
ok  	github.com/sunilgopinath/algomap/hashmaps/1189NumberOfBalloons	(cached)
ok  	github.com/sunilgopinath/algomap/hashmaps/128LongestSubsequence	(cached)
ok  	github.com/sunilgopinath/algomap/hashmaps/169MajorityElement	(cached)
ok  	github.com/sunilgopinath/algomap/hashmaps/1TwoSum	(cached)
ok  	github.com/sunilgopinath/algomap/hashmaps/217ContainsDuplicates	(cached)
ok  	github.com/sunilgopinath/algomap/hashmaps/242ValidAnagram	(cached)
ok  	github.com/sunilgopinath/algomap/hashmaps/36ValidSudoku	(cached)
ok  	github.com/sunilgopinath/algomap/hashmaps/383RansomNote	(cached)
ok  	github.com/sunilgopinath/algomap/hashmaps/49GroupAnagrams	(cached)
ok  	github.com/sunilgopinath/algomap/hashmaps/771JewelsandStones	(cached)
ok  	github.com/sunilgopinath/algomap/linkedLists/83RemoveDuplicates	(cached)
ok  	github.com/sunilgopinath/algomap/stacks/150ReversePolishNotation	(cached)
ok  	github.com/sunilgopinath/algomap/stacks/20Parenthesis	(cached)
?   	github.com/sunilgopinath/algomap/stacks/739DailyTemperatures	[no test files]
ok  	github.com/sunilgopinath/algomap/stacks/682BaseballGame	(cached)
ok  	github.com/sunilgopinath/algomap/twoPointers/167TwoSumII	0.513s
ok  	github.com/sunilgopinath/algomap/twoPointers/344ReverseString	(cached)
ok  	github.com/sunilgopinath/algomap/twoPointers/977SquaresOfSortedArrays	(cached)
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