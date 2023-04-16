package main

import "fmt"

// It can be proved that the first missing positive integer is a number up to N+1
// We use the array itself as a memory to save occurance of each number between 1 and N (0 index refers to N)
func solve(arr []int) int {
	// first check for 1
	for i := 0; i < len(arr); i++ {
		if !OneExists(arr) {
			return 1
		}
	}

	// clean array
	for i, v := range arr {
		if v < 1 || v > len(arr) {
			arr[i] = 1
		}
	}

	// mark data
	for _, v := range arr {
		if localAbs(v) == len(arr) && arr[0] > 0 {
			arr[0] *= -1
		} else if arr[localAbs(v)] > 0 {
			arr[localAbs(v)] *= -1
		}
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] > 0 {
			return i
		}
	}

	if arr[0] > 0 {
		return len(arr)
	}

	return len(arr) + 1
}

func OneExists(arr []int) bool {
	for _, v := range arr {
		if v == 1 {
			return true
		}
	}
	return false
}

func localAbs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

// input1: [3, 4, -1, 1]
// output1: 2
// input2: [1, 2, 0]
// output2: 3
func main() {
	arr1 := []int{3, 4, -1, 1}
	fmt.Println(solve(arr1))

	arr2 := []int{1, 2, 0}
	fmt.Println(solve(arr2))
}
