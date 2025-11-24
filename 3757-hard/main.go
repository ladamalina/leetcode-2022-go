package main

const MOD int64 = 1e9 + 7

func countEffective(nums []int) int {
	pow2 := make([]int64, len(nums)+1)
	pow2[0] = 1
	for i := 1; i <= len(nums); i++ {
		pow2[i] = (pow2[i-1] * 2) % MOD
	}
	fullOr := 0
	for _, x := range nums {
		fullOr |= x
	}
	nMasks := fullOr + 1
	dp := make([]int64, nMasks)
	for _, x := range nums {
		dp[x]++
	}
	for b := 1; b < nMasks; b <<= 1 {
		for mask := 0; mask < nMasks; mask++ {
			if (mask & b) != 0 {
				dp[mask] += dp[mask^b]
			}
		}
	}
	for mask := 0; mask < nMasks; mask++ {
		dp[mask] = (pow2[dp[mask]] - 1 + MOD) % MOD
	}
	for b := 1; b < nMasks; b <<= 1 {
		for mask := 0; mask < nMasks; mask++ {
			if (mask & b) != 0 {
				dp[mask] = (dp[mask] - dp[mask^b] + MOD) % MOD
			}
		}
	}

	var res int64 = 1
	for mask := 0; mask < fullOr; mask++ {
		res = (res + dp[mask]) % MOD
	}
	return int(res)
}
