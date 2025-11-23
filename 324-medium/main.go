package main

import (
	"slices"
)

func wiggleSort(nums []int) {
	slices.Sort(nums)
	arr := make([]int, len(nums))
	l, r := (len(nums)-1)>>1, len(nums)-1
	useL := true
	for i := 0; i < len(nums); i++ {
		if useL {
			arr[i] = nums[l]
			l--
		} else {
			arr[i] = nums[r]
			r--
		}
		useL = !useL
	}
	copy(nums, arr)
}
