package main

import (
	"strconv"
)

const maxD = 10
const maxN = 16

var dpC [2][3][maxD][maxD][maxN]int64
var dpW [2][3][maxD][maxD][maxN]int64

func rec(t, z, d1, d2, i int, s *string) (int64, int64) {
	if i == len(*s) {
		return 1, 0
	}
	resC, resW := &dpC[t][z][d1][d2][i], &dpW[t][z][d1][d2][i]
	if *resC != -1 {
		return *resC, *resW
	}
	*resC, *resW = 0, 0
	mn, mx, sd := 0, 9, int((*s)[i]-'0')
	if t == 1 {
		mx = sd
	}
	for d3 := mn; d3 <= mx; d3++ {
		var add int64 = 0
		if z == 2 {
			if (d1 < d2 && d2 > d3) || (d1 > d2 && d2 < d3) {
				add = 1
			}
		}
		nt := t
		if d3 < sd {
			nt = 0
		}
		nz := z
		if z == 0 && d3 > 0 {
			nz = 1
		}
		if z == 1 {
			nz = 2
		}
		subC, subW := rec(nt, nz, d2, d3, i+1, s)
		*resC += subC
		*resW += add*subC + subW
	}
	return *resC, *resW
}

func wavinessLte(num int64) int64 {
	s := strconv.FormatInt(num, 10)
	if len(s) < 3 {
		return 0
	}
	for t := 0; t < 2; t++ {
		for z := 0; z < 3; z++ {
			for d1 := 0; d1 < maxD; d1++ {
				for d2 := 0; d2 < maxD; d2++ {
					for i := 0; i < len(s); i++ {
						dpC[t][z][d1][d2][i] = -1
						dpW[t][z][d1][d2][i] = -1
					}
				}
			}
		}
	}
	_, w := rec(1, 0, 0, 0, 0, &s)
	return w
}

func waviness(num int64) int64 {
	s := strconv.FormatInt(num, 10)
	var res int64 = 0
	for i := 1; i < len(s)-1; i++ {
		if (s[i-1] < s[i] && s[i] > s[i+1]) || (s[i-1] > s[i] && s[i] < s[i+1]) {
			res++
		}
	}
	return res
}

func totalWaviness(num1 int64, num2 int64) int64 {
	//logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
	//	Level: slog.LevelDebug,
	//}))

	w1 := wavinessLte(num1)
	w2 := wavinessLte(num2)
	w3 := waviness(num1)
	res := w2 - w1 + w3
	//logger.Debug("", "num1", num1, "wavinessLte", w1, "waviness", w3)
	//logger.Debug("", "num2", num2, "wavinessLte", w2)
	//logger.Debug("", "res", res)
	return res
}
