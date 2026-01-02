package main

func countCollisions(s string) int {
	res := 0
	var hasRS, hasLS bool
	for _, c := range s {
		if c == 'L' && hasRS {
			res++
		}
		hasRS = hasRS || c == 'R' || c == 'S'
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'R' && hasLS {
			res++
		}
		hasLS = hasLS || s[i] == 'L' || s[i] == 'S'
	}
	return res
}
