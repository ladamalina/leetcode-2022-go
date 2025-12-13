package main

func maxSubgraphScore(n int, edges [][]int, good []int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	sumDown := make([]int, n)
	var dfsMaxDown func(u, parent int)
	dfsMaxDown = func(u, parent int) {
		if good[u] == 1 {
			sumDown[u] = 1
		} else {
			sumDown[u] = -1
		}
		for _, v := range g[u] {
			if v == parent {
				continue
			}
			dfsMaxDown(v, u)
			if sumDown[v] > 0 {
				sumDown[u] += sumDown[v]
			}
		}
	}
	dfsMaxDown(0, -1)

	res := make([]int, n)

	var dfsRes func(u, parent, parentContrib int)
	dfsRes = func(u, parent, parentContrib int) {
		res[u] = sumDown[u] + parentContrib
		for _, v := range g[u] {
			if v == parent {
				continue
			}
			uContrib := res[u]
			if sumDown[v] > 0 {
				uContrib -= sumDown[v]
			}
			dfsRes(v, u, max(0, uContrib))
		}
	}
	dfsRes(0, -1, 0)

	return res
}
