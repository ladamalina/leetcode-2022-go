package main

const maxK int = 1e5

var used [maxK]bool

func smallestRepunitDivByK(k int) int {
	for i := 1; i < maxK; i++ {
		used[i] = false
	}
	mod, n := 1%k, 1
	used[mod] = true
	for mod != 0 {
		mod = (mod*10 + 1) % k
		n++
		if mod == 0 {
			return n
		}
		if used[mod] {
			return -1
		}
		used[mod] = true
	}
	return n
}
