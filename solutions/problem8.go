package main

import "fmt"

// global var to count Unival Trees
var c int

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func isUnivalTree(node *Node) (bool, int) {
	lvalid, rvalid := true, true
	lvalue, rvalue := node.Val, node.Val
	if node.Left != nil {
		lvalid, lvalue = isUnivalTree(node.Left)
	}
	if node.Right != nil {
		rvalid, rvalue = isUnivalTree(node.Right)
	}

	if !lvalid || lvalue != node.Val || !rvalid || rvalue != node.Val {
		return false, 0
	}

	c += 1
	return true, node.Val
}

func main() {
	root := &Node{0,
		&Node{1, nil, nil},
		&Node{0, &Node{1, &Node{1, nil, nil}, &Node{1, nil, nil}}, &Node{0, nil, nil}}}

	isUnivalTree(root)

	fmt.Println(c)
}
