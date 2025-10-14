package main

func getMsb(x int64) int {
	for b := 63; b >= 0; b-- {
		if x&(1<<b) != 0 {
			return b
		}
	}
	return -1
}

func makePal(pref int, len int) int64 {
	pal := int64(pref)
	q := int64(pref)
	if len&1 == 1 {
		q >>= 1
	}
	for q > 0 {
		pal = (pal << 1) | (q & 1)
		q >>= 1
	}
	return pal
}

const maxNBits = 50

var dp [maxNBits + 1]int

func countBinaryPalindromes(n int64) int {
	msb := getMsb(n)
	if msb == -1 {
		return 1
	}
	nBits := msb + 1

	dp[1], dp[2] = 1, 1
	for i := 3; i <= nBits; i++ {
		dp[i] = dp[i-2] << 1
	}

	cnt := 1 // 0
	for i := 1; i < nBits; i++ {
		cnt += dp[i]
	}

	hBits := (nBits + 1) >> 1
	minMask, maxMask := 1<<(hBits-1), (1<<hBits)-1
	lo, hi := minMask, maxMask
	for lo < hi {
		mid := (lo + hi + 1) >> 1
		pal := makePal(mid, nBits)
		if pal <= n {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	if makePal(hi, nBits) <= n {
		cnt += hi - minMask + 1
	}

	return cnt
}
