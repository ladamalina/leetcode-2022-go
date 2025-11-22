package main

import (
	"slices"
)

func wiggleSort(nums []int) {
	slices.Sort(nums)
	arr := make([]int, len(nums))
	i, l, r := 0, (len(nums)-1)>>1, len(nums)-1
	use_l := true
	for i < len(nums) {
		if use_l {
			arr[i] = nums[l]
			l--
		} else {
			arr[i] = nums[r]
			r--
		}
		use_l = !use_l
		i++
	}
	copy(nums, arr)
}
