package main

import "sort"

func searchMatrix(g [][]int, target int) bool {
	if len(g) == 0 || len(g[0]) == 0 {
		return false
	}

	cut := false

	n := sort.Search(len(g[0]), func(i int) bool {
		return g[0][i] >= target
	})
	if n < len(g[0]) {
		if g[0][n] == target {
			return true
		}
		if n == 0 {
			return false
		}
		for i := range g {
			g[i] = g[i][0:n]
		}
		cut = true
	}

	n = sort.Search(len(g[0]), func(i int) bool {
		return g[len(g)-1][i] >= target
	})
	if n == len(g[0]) {
		return false
	}
	if g[len(g)-1][n] == target {
		return true
	}
	if n > 0 {
		for i := range g {
			g[i] = g[i][n:]
		}
		cut = true
	}

	n = sort.Search(len(g), func(i int) bool {
		return g[i][0] >= target
	})
	if n < len(g) {
		if g[n][0] == target {
			return true
		}
		if n == 0 {
			return false
		}
		g = g[0:n]
	}

	n = sort.Search(len(g), func(i int) bool {
		return g[i][len(g[i])-1] >= target
	})
	if n == len(g) {
		return false
	}
	if g[n][len(g[n])-1] == target {
		return true
	}
	if n > 0 {
		g = g[n:]
		cut = true
	}

	if !cut {
		return false
	}

	return searchMatrix(g, target)
}
