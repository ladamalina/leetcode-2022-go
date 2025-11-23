package main

func longestSubarray(nums []int) int {
	maxRes, res := 2, 2
	for i, x := range nums {
		if i < 2 {
			continue
		}
		if x == nums[i-2]+nums[i-1] {
			res++
			maxRes = max(maxRes, res)
		} else {
			res = 2
		}
	}
	return maxRes
}
