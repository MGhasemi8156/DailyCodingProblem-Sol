package main

import "fmt"

func solve1(arr []int) []int {
	all := 1
	for _, v := range arr {
		all *= v
	}

	res := make([]int, len(arr))
	for i, _ := range res {
		res[i] = all / arr[i]
	}
	return res
}

// solution regarding follow-up
func solve2(arr []int) []int {
	n := len(arr)
	prefix := make([]int, n)
	suffix := make([]int, n)

	prefix[0] = arr[0]
	for i := 1; i < n; i++ {
		prefix[i] = arr[i] * prefix[i-1]
	}
	suffix[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = arr[i] * suffix[i+1]
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		a, b := 1, 1
		if i != 0 {
			a = prefix[i-1]
		}
		if i != n-1 {
			b = suffix[i+1]
		}
		res[i] = a * b
	}

	return res
}

// sample: arr = [1, 2, 3, 4, 5]

func main() {
	arr := []int{1, 2, 3, 4, 5}

	res1 := solve1(arr)
	fmt.Println(res1)

	res2 := solve2(arr)
	fmt.Println(res2)
}
