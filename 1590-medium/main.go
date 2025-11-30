package main

func minSubarray(nums []int, p int) int {
	sumMod := 0
	for _, x := range nums {
		sumMod = (sumMod + x) % p
	}
	if sumMod == 0 {
		return 0
	}

	modIdx := make(map[int]int)
	modIdx[0] = -1
	minRes := len(nums)
	mod := 0
	for i := 0; i < len(nums); i++ {
		mod = (mod + nums[i]) % p
		prevMod := (mod - sumMod + p) % p
		j, found := modIdx[prevMod]
		if found {
			minRes = min(minRes, i-j)
		}
		modIdx[mod] = i
	}
	if minRes == len(nums) {
		return -1
	}
	return minRes
}
