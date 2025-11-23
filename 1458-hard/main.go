package main

const maxN int = 500
const inf int = -1e9

var dp [2][maxN][maxN]int

func rec(i, j, k int, nums1 []int, nums2 []int) int {
	if j == len(nums1) || k == len(nums2) {
		if i == 0 {
			return inf
		}
		return 0
	}
	res := &dp[i][j][k]
	if *res != inf {
		return *res
	}

	// take pair
	*res = max(*res, nums1[j]*nums2[k]+rec(1, j+1, k+1, nums1, nums2))
	// skip nums1[j]
	*res = max(*res, rec(i, j+1, k, nums1, nums2))
	// skip nums2[k]
	*res = max(*res, rec(i, j, k+1, nums1, nums2))

	return *res
}

func maxDotProduct(nums1 []int, nums2 []int) int {
	for i := 0; i < 2; i++ {
		for j := 0; j < len(nums1); j++ {
			for k := 0; k < len(nums2); k++ {
				dp[i][j][k] = inf
			}
		}
	}
	return rec(0, 0, 0, nums1, nums2)
}
