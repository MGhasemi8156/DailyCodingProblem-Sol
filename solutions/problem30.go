package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(arr []int) int {
	right := make([]int, len(arr))
	left := make([]int, len(arr))

	leftMax := 0
	for i := 0; i < len(arr); i++ {
		left[i] = leftMax
		leftMax = max(arr[i], leftMax)
	}

	rightMax := 0
	for i := len(arr) - 1; i >= 0; i-- {
		right[i] = rightMax
		rightMax = max(arr[i], rightMax)
	}

	result := 0
	for i := 0; i < len(arr); i++ {
		result += max(arr[i]-min(right[i], left[i]), 0)
	}
	return result
}

// input: [2, 1, 2]
// output: 1
// input: [3, 0, 1, 3, 0, 5]
// output: 8

func main() {
	input1 := []int{2, 1, 2}
	fmt.Println(solve(input1))

	input2 := []int{3, 0, 1, 3, 0, 5}
	fmt.Println(solve(input2))
}
