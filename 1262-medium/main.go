package main

import (
	"slices"
)

func maxSumDivThree(nums []int) int {
	modVals := make([][]int, 3)
	for _, x := range nums {
		modVals[x%3] = append(modVals[x%3], x)
	}
	s := 0
	for mod := 0; mod < 3; mod++ {
		for _, x := range modVals[mod] {
			s += x
		}
	}
	slices.Sort(modVals[1])
	slices.Sort(modVals[2])
	if s%3 == 2 {
		s11, s2 := -1, -1
		if len(modVals[1]) >= 2 {
			s11 = s - modVals[1][0] - modVals[1][1]
		}
		if len(modVals[2]) >= 1 {
			s2 = s - modVals[2][0]
		}
		return max(s11, s2)
	}
	if s%3 == 1 {
		s1, s22 := -1, -1
		if len(modVals[1]) >= 1 {
			s1 = s - modVals[1][0]
		}
		if len(modVals[2]) >= 2 {
			s22 = s - modVals[2][0] - modVals[2][1]
		}
		return max(s1, s22)
	}
	// maxRes%3 == 0
	return s
}
