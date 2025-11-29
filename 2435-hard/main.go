package main

const MOD int64 = 1e9 + 7

func numberOfPaths(g [][]int, k int) int {
	m, n := len(g), len(g[0])
	dp := make([][][]int64, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int64, k)
		}
	}
	dp[m-1][n-1][g[m-1][n-1]%k] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}
			for pk := 0; pk < k; pk++ {
				ck := (g[i][j] + pk) % k
				if i+1 < m {
					dp[i][j][ck] = (dp[i][j][ck] + dp[i+1][j][pk]) % MOD
				}
				if j+1 < n {
					dp[i][j][ck] = (dp[i][j][ck] + dp[i][j+1][pk]) % MOD
				}
			}
		}
	}
	return int(dp[0][0][0])
}
