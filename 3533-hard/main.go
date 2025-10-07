package main

import (
	"slices"
	"strconv"
)

func concatenatedDivisibility(nums []int, k int) []int {
	maxMask := (1 << len(nums)) - 1
	dp := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, maxMask+1)
		for j := 0; j <= maxMask; j++ {
			dp[i][j] = -1
		}
	}
	slices.Sort(nums)
	sumLen := 0
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
		sumLen += len(strs[i])
	}

	resArr := make([]int, 0)

	mult := func(mod int, s *string) int {
		for i := 0; i < len(*s); i++ {
			mod = (mod*10 + int((*s)[i]-'0')) % k
		}
		return mod
	}

	var rec func(mask, mod, l int) int
	rec = func(mask, mod, l int) int {
		if mask == maxMask {
			if mod == 0 {
				return 1
			}
			return 0
		}
		res := &dp[mod][mask]
		if *res != -1 {
			return *res
		}
		*res = 0
		for b := 0; b < len(nums); b++ {
			if (mask & (1 << b)) == 0 {
				resArr = append(resArr, nums[b])
				subRes := rec(mask|(1<<b), mult(mod, &strs[b]), l-len(strs[b]))
				if subRes == 1 {
					*res = 1
					break
				}
				resArr = resArr[:len(resArr)-1]
			}
		}
		return *res
	}
	rec(0, 0, sumLen)
	return resArr
}
