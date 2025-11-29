package main

func maxSubarraySum(nums []int, k int) int64 {
	dp := make([]int64, k)
	for i := 0; i < k; i++ {
		dp[i] = 2e18
	}
	var psum, maxRes int64 = 0, -2e18
	dp[0] = psum
	for i := 0; i < len(nums); i++ {
		psum += int64(nums[i])
		if i >= k-1 {
			maxRes = max(maxRes, psum-dp[(i+1)%k])
		}
		dp[(i+1)%k] = min(dp[(i+1)%k], psum)
	}
	return maxRes
}
