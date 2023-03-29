package main

import "fmt"

func solve(arr []int, k int) []int {
	nums := make(map[int]bool)

	for _, v := range arr {
		nums[v] = true
		_, exists := nums[k-v]
		if exists {
			return []int{v, k - v}
		}
	}
	return nil
}

// sample: arr = [10, 15, 3, 7] k = 17

func main() {
	arr := []int{10, 15, 3, 7}
	k := 17

	res := solve(arr, k)

	fmt.Println(res[0], "+", res[1], "=", k)
}
