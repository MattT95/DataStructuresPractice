package main

import (
	"fmt"
)

func main() {

	tree := Node{
		0,
		&Node{1, nil, nil},
		&Node{
			0,
			&Node{
				1,
				&Node{1, nil, nil},
				&Node{1, nil, nil},
			},
			&Node{
				0,
				nil,
				nil,
			},
		},
	}

	count := CountUnivalTrees(&tree)

	fmt.Printf("Expecting %d. Got %d", 5, count)
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func CountUnivalTrees(node *Node) int {

	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	if IsUnivalTree(node, node.Value) {
		return 1 + CountUnivalTrees(node.Left) + CountUnivalTrees(node.Right)
	} else {
		return CountUnivalTrees(node.Left) + CountUnivalTrees(node.Right)
	}
}

func IsUnivalTree(node *Node, value int) bool {

	if node == nil || node.Value != value {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return true
	}

	return IsUnivalTree(node.Left, value) && IsUnivalTree(node.Right, value)
}
