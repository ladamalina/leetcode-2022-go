package main

func rec(nums *[]int, x, l, r int) bool {
	if l == r {
		return (*nums)[l] == x
	}
	m := (l + r) >> 1
	if (*nums)[m] == x {
		return true
	} // x != nums[m]
	if (*nums)[l] < (*nums)[r] { // monotonic
		if x < (*nums)[l] || x > (*nums)[r] {
			return false
		}
		if x < (*nums)[m] {
			return rec(nums, x, l, m)
		}
		return rec(nums, x, m+1, r)
	} // shift inside
	return rec(nums, x, l, m) || rec(nums, x, m+1, r)
}

func search(nums []int, target int) bool {
	return rec(&nums, target, 0, len(nums)-1)
}
