package main

import (
	"strconv"
)

func waviness(num int) int {
	s := strconv.FormatInt(int64(num), 10)
	var res int = 0
	for i := 1; i < len(s)-1; i++ {
		if (s[i-1] < s[i] && s[i] > s[i+1]) || (s[i-1] > s[i] && s[i] < s[i+1]) {
			res++
		}
	}
	return res
}

func totalWaviness(num1 int, num2 int) int {
	res := 0
	for x := num1; x <= num2; x++ {
		res += waviness(x)
	}
	return res
}
