package main

import "fmt"

func solve(n int) int {
	save := make([]int, n+1)

	var sol func(i int) int
	sol = func(i int) int {
		if save[i] != 0 {
			return save[i]
		}
		if i == 1 {
			return 1
		}
		if i == 2 {
			return 2
		}
		return sol(i-1) + sol(i-2)
	}
	return sol(n)
}

// for 1, 3, 5 step size is the same with different base conditions

// input: 4
// output: 5
func main() {
	fmt.Println(solve(4))
}
