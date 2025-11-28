package main

const MOD int64 = 1e9 + 7
const maxN, maxG = 100, 100

var dp [maxG + 1][maxN + 1]int64

func numMusicPlaylists(n int, goal int, k int) int {
	for i := 0; i <= goal; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = 0
		}
	}
	dp[0][0] = 1
	for i := 1; i <= goal; i++ {
		for j := 1; j <= min(i, n); j++ {
			addNew := (dp[i-1][j-1] * int64(n-j+1)) % MOD
			dp[i][j] = (dp[i][j] + addNew) % MOD
			if k < j {
				addRep := (dp[i-1][j] * int64(j-k)) % MOD
				dp[i][j] = (dp[i][j] + addRep) % MOD
			}
		}
	}
	return int(dp[goal][n])
}
