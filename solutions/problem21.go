package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func solve(intervals []Interval) int {
	n := len(intervals)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, v := range intervals {
		starts[i] = v.Start
		ends[i] = v.End
	}

	sort.Ints(starts)
	sort.Ints(ends)

	count, res := 0, 0
	for i, j := 0, 0; i < n || j < n; {
		if i == n || ends[j] <= starts[i] {
			count--
			j++
		} else {
			count++
			if count > res {
				res = count
			}
			i++
		}
	}
	return res
}

// input: [(30, 75), (0, 50), (60, 150)]
// output: 2
func main() {
	intervals := []Interval{{30, 75}, {0, 50}, {60, 150}}

	fmt.Println(solve(intervals))
}
