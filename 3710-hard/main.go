package main

//import (
//	"log/slog"
//	"os"
//)

const maxN = 500

var d [maxN][maxN]int
var vis [maxN]int

func dist(p1, p2 []int) int {
	var dx, dy int
	if p1[0] >= p2[0] {
		dx = p1[0] - p2[0]
	} else {
		dx = p2[0] - p1[0]
	}
	if p1[1] >= p2[1] {
		dy = p1[1] - p2[1]
	} else {
		dy = p2[1] - p1[1]
	}
	return dx + dy
}

func binCheck(n, lim int) bool {
	for i := 0; i < n; i++ {
		vis[i] = -1
	}
	cnt := make([]int, 2)
	nextC := 1

	var dfs func(int, int) bool
	dfs = func(u, c int) bool {
		if vis[u] == 1-c {
			return false
		}
		if vis[u] == c {
			return true
		}
		vis[u] = c
		cnt[c]++
		for v := 0; v < n; v++ {
			if v != u && d[u][v] < lim && !dfs(v, 1-c) {
				return false
			}
		}
		return true
	}

	for v0 := 0; v0 < n; v0++ {
		if vis[v0] == -1 {
			if !dfs(v0, nextC) {
				return false
			}
			nextC = 1 - nextC
		}
	}
	if max(cnt[0], cnt[1]) == 1 {
		return lim == 0
	}
	return true
}

func maxPartitionFactor(points [][]int) int {
	//logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
	//	Level: slog.LevelDebug,
	//}))

	n := len(points)
	maxD := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			d[i][j] = dist(points[i], points[j])
			maxD = max(maxD, d[i][j])
		}
	}
	lo, hi := 0, maxD
	for lo < hi {
		mid := (lo + hi + 1) >> 1
		check := binCheck(n, mid)
		//logger.Debug("", "lo", lo, "mid", mid, "hi", hi, "check", check)
		if check {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return hi
}
