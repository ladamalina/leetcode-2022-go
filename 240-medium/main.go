package main

import "sort"

func search(g *[][]int, target int) bool {
	if len(*g) == 0 {
		return false
	}

	if len((*g)[0]) == 0 {
		return false
	}

	cut := false

	n1 := sort.Search(len((*g)[0]), func(i int) bool {
		return (*g)[0][i] >= target
	})
	if n1 < len((*g)[0]) {
		if (*g)[0][n1] == target {
			return true
		}
		if n1 == 0 {
			return false
		}
		for i := range *g {
			(*g)[i] = (*g)[i][0:n1]
		}
		cut = true
	}

	n2 := sort.Search(len((*g)[0]), func(i int) bool {
		return (*g)[len(*g)-1][i] >= target
	})
	if n2 == len((*g)[0]) {
		return false
	}
	if (*g)[len(*g)-1][n2] == target {
		return true
	}
	if n2 > 0 {
		for i := range *g {
			(*g)[i] = (*g)[i][n2:]
		}
		cut = true
	}

	n3 := sort.Search(len(*g), func(i int) bool {
		return (*g)[i][0] >= target
	})
	if n3 < len(*g) {
		if (*g)[n3][0] == target {
			return true
		}
		if n3 == 0 {
			return false
		}
		*g = (*g)[0:n3]
	}

	n4 := sort.Search(len(*g), func(i int) bool {
		return (*g)[i][len((*g)[i])-1] >= target
	})
	if n4 == len(*g) {
		return false
	}
	if (*g)[n4][len((*g)[n4])-1] == target {
		return true
	}
	if n4 > 0 {
		*g = (*g)[n4:]
		cut = true
	}

	if !cut {
		return false
	}

	return search(g, target)
}

func searchMatrix(g [][]int, target int) bool {
	return search(&g, target)
}
