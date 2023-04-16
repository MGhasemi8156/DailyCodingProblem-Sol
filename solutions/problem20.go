package main

import "fmt"

type Node struct {
	Val  int
	next *Node
}

func findIntersectionPoint(head1 *Node, head2 *Node) *Node {
	current1 := head1
	current2 := head2

	for current1 != nil && current2 != nil && current1 != current2 {
		current1 = current1.next
		current2 = current2.next

		if current1 == nil {
			current1 = head2
		}
		if current2 == nil {
			current2 = head1
		}
	}

	return current1
}

// input  A = 3 -> 7 -> 8 -> 10 and B = 99 -> 1 -> 8 -> 10
// output 8 (node)

func main() {
	intersectionPoint := &Node{8, &Node{10, nil}}
	head1 := &Node{3, &Node{7, intersectionPoint}}
	head2 := &Node{99, &Node{1, intersectionPoint}}

	result := findIntersectionPoint(head1, head2)
	fmt.Println(result.Val)
}
