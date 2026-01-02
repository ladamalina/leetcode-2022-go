package main

func bestClosingTime(customers string) int {
	n := len(customers)
	cntY := 0
	for i := 0; i < n; i++ {
		if customers[i] == 'Y' {
			cntY++
		}
	}
	bestPenalty, bestDay := n-cntY, n
	befN, aftY := 0, cntY
	for i := 0; i < n; i++ {
		penalty := befN + aftY
		if penalty < bestPenalty || (penalty == bestPenalty && i < bestDay) {
			bestPenalty = penalty
			bestDay = i
		}
		if customers[i] == 'Y' {
			aftY--
		} else {
			befN++
		}
	}
	return bestDay
}
