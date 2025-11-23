package main

import (
	"slices"
	"sort"
)

const maxN int = 1e5

var dp [2][maxN]int

type Robot struct {
	x, d                   int
	minL, maxL, minR, maxR int
}

var rs [maxN]Robot

func maxWalls(robots []int, distance []int, walls []int) int {
	//logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
	//	Level: slog.LevelDebug,
	//}))
	n, m := len(robots), len(walls)

	for i := 0; i < n; i++ {
		rs[i] = Robot{robots[i], distance[i], m, -1, m, -1}
	}
	slices.SortFunc(rs[:n], func(a, b Robot) int {
		if a.x < b.x {
			return -1
		}
		if a.x > b.x {
			return 1
		}
		return 0
	})
	slices.Sort(walls)

	for i := 0; i < n; i++ {
		lx := rs[i].x - rs[i].d
		if i > 0 {
			lx = max(lx, rs[i-1].x)
		}
		rx := rs[i].x + rs[i].d
		if i+1 < n {
			rx = min(rx, rs[i+1].x)
		}

		lLo := sort.Search(m, func(j int) bool {
			return walls[j] >= lx
		})
		lUp := sort.Search(m, func(j int) bool {
			return walls[j] > rs[i].x
		})
		if lLo < lUp {
			rs[i].minL, rs[i].maxL = lLo, lUp-1
		}

		rLo := sort.Search(m, func(j int) bool {
			return walls[j] >= rs[i].x
		})
		rUp := sort.Search(m, func(j int) bool {
			return walls[j] > rx
		})
		if rLo < rUp {
			rs[i].minR, rs[i].maxR = rLo, rUp-1
		}
	}

	//for i := 0; i < n; i++ {
	//	logger.Debug("Robot", "x", rs[i].x, "d", rs[i].d)
	//	if rs[i].minL <= rs[i].maxL {
	//		logger.Debug("lWalls", "", walls[rs[i].minL:rs[i].maxL+1])
	//	}
	//	if rs[i].minR <= rs[i].maxR {
	//		logger.Debug("rWalls", "", walls[rs[i].minR:rs[i].maxR+1])
	//	}
	//}

	for dir := 0; dir < 2; dir++ {
		for i := 0; i < n; i++ {
			dp[dir][i] = -1
		}
	}

	var rec func(dir, i int) int
	rec = func(dir, i int) int {
		if i == n {
			return 0
		}
		res := &dp[dir][i]
		if *res != -1 {
			return *res
		}
		*res = 0

		lCnt, rCnt := 0, 0
		if dir == 0 {
			minL := rs[i].minL
			if i > 0 {
				minL = max(rs[i-1].maxL+1, rs[i].minL)
			}
			if minL <= rs[i].maxL {
				lCnt = rs[i].maxL - minL + 1
			}
			if rs[i].minR <= rs[i].maxR {
				rCnt = rs[i].maxR - rs[i].minR + 1
			}
		} else {
			minL := max(rs[i-1].maxR+1, rs[i].minL)
			if minL <= rs[i].maxL {
				lCnt = rs[i].maxL - minL + 1
			}
			minR := max(rs[i-1].maxR+1, rs[i].minR)
			if minR <= rs[i].maxR {
				rCnt = rs[i].maxR - minR + 1
			}
		}
		*res = max(*res, lCnt+rec(0, i+1))
		*res = max(*res, rCnt+rec(1, i+1))
		return *res
	}
	res := rec(0, 0)

	return res
}
