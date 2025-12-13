package main

func tilingRectangle(n int, m int) int {
	if n == m {
		return 1
	}
	if n > m {
		n, m = m, n
	}
	heights := make([]int, n)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i == j {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = i * j
			for k := 1; k <= i/2; k++ {
				dp[i][j] = min(dp[i][j], dp[i-k][j]+dp[k][j])
			}
			for k := 1; k <= j/2; k++ {
				dp[i][j] = min(dp[i][j], dp[i][j-k]+dp[i][k])
			}
		}
	}

	res := dp[n][m]
	dfs(n, m, heights, 0, &res)
	return res
}

func dfs(n, m int, heights []int, cnt int, res *int) {
	if cnt >= *res {
		return
	}

	left, min_height := 0, m
	for i := 0; i < n; i++ {
		if heights[i] < min_height {
			left = i
			min_height = heights[i]
		}
	}

	if min_height == m {
		*res = min(cnt, *res)
		return
	}

	right := left
	for right < n && heights[left] == heights[right] && right-left+min_height < m {
		right++
	}

	for i := left; i < right; i++ {
		heights[i] += right - left
	}

	for size := right - left; size >= 1; size-- {
		dfs(n, m, heights, cnt+1, res)

		// update the rectangle to contain the next smaller square
		for i := left; i < left+size-1; i++ {
			heights[i]--
		}
		heights[left+size-1] -= size
	}
}
