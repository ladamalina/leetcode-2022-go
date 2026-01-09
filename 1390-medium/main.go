package main

func checkDivisors(x int) (int, int) {
	cnt, sum := 0, 0
	for f1 := 1; f1*f1 <= x; f1++ {
		if x%f1 == 0 {
			f2 := x / f1
			if f1 == f2 {
				cnt++
				sum += f1
			} else {
				cnt += 2
				sum += f1 + f2
			}
		}
	}
	return cnt, sum
}

func sumFourDivisors(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		cnt, sum := checkDivisors(nums[i])
		if cnt == 4 {
			res += sum
		}
	}
	return res
}
