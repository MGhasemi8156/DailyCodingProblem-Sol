package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func solve(head *Node, k int) *Node {
	result := head
	for i := 0; i < k+1; i++ {
		head = head.Next
	}

	for head != nil {
		head = head.Next
		result = result.Next
	}

	return result
}

func main() {
	head := &Node{111, &Node{222, &Node{333, &Node{444, &Node{555, &Node{666, &Node{777, nil}}}}}}}

	fmt.Println(solve(head, 2).Val)
}
