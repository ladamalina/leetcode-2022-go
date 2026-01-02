package main

import "math/bits"

const mod = 1e9 + 7

type SparseTable struct {
	rmqMin [][]int
	rmqMax [][]int
}

func NewSparseTable(a []int) *SparseTable {
	n := len(a)
	if n == 0 {
		return &SparseTable{}
	}

	nBits := bits.Len(uint(n))

	rmqMin := make([][]int, nBits)
	rmqMax := make([][]int, nBits)

	rmqMin[0] = make([]int, n)
	rmqMax[0] = make([]int, n)
	copy(rmqMin[0], a)
	copy(rmqMax[0], a)

	for h := 1; h < nBits; h++ {
		length := n - (1 << h) + 1
		if length < 0 {
			length = 0
		}
		rmqMin[h] = make([]int, length)
		rmqMax[h] = make([]int, length)

		half := 1 << (h - 1)
		for i := 0; i < length; i++ {
			rmqMin[h][i] = min(rmqMin[h-1][i], rmqMin[h-1][i+half])
			rmqMax[h][i] = max(rmqMax[h-1][i], rmqMax[h-1][i+half])
		}
	}

	return &SparseTable{
		rmqMin: rmqMin,
		rmqMax: rmqMax,
	}
}

func (st *SparseTable) GetMin(l, r int) int {
	length := r - l + 1
	k := bits.Len(uint(length)) - 1
	return min(
		st.rmqMin[k][l],
		st.rmqMin[k][r-(1<<k)+1],
	)
}

func (st *SparseTable) GetMax(l, r int) int {
	length := r - l + 1
	k := bits.Len(uint(length)) - 1
	return max(
		st.rmqMax[k][l],
		st.rmqMax[k][r-(1<<k)+1],
	)
}

func countPartitions(nums []int, k int) int {
	dp := make([]int, len(nums)+2)
	dp[len(nums)] = 1
	dp[len(nums)+1] = 0

	dp_ssum := make([]int, len(nums)+2)
	dp_ssum[len(nums)] = 1
	dp_ssum[len(nums)+1] = 0

	st := NewSparseTable(nums)

	for i := len(nums) - 1; i >= 0; i-- {
		lo, hi := i, len(nums)-1
		for lo < hi {
			mid := (lo + hi + 1) >> 1
			mn := st.GetMin(i, mid)
			mx := st.GetMax(i, mid)
			if mx-mn <= k {
				lo = mid
			} else {
				hi = mid - 1
			}
		}
		dp[i] = (dp_ssum[i+1] - dp_ssum[hi+2] + mod) % mod
		dp_ssum[i] = (dp_ssum[i+1] + dp[i]) % mod
	}
	return dp[0]
}
