package main

import "fmt"

// Find the 2nd largest node in a binary search tree

func main() {

	tree := Node{
		32,
		&Node{
			16,
			&Node{
				8,
				&Node{4, nil, nil},
				nil,
			},
			&Node{
				24,
				&Node{22, nil, nil},
				&Node{26, nil, nil},
			},
		},
		&Node{
			64,
			&Node{48, nil, nil},
			&Node{72, nil, nil},
		},
	}

	secondLargest := FindSecondLargest(&tree, -1)

	fmt.Println(secondLargest)
}

type Node struct {
	value int
	Left  *Node
	Right *Node
}

func FindSecondLargest(root *Node, previousNodeValue int) int {

	// If the right isn't empty, there is at least one larger node than the current,
	// So we need to traverse
	if root.Right != nil {
		return FindSecondLargest(root.Right, root.value)
	}

	// If the left isn't empty, but the right is, then this is the largest node
	// and we have at least one node that
	// is both smaller than this one and larger than the previous in the sub tree
	// and therefore a candidate for the second largest
	// We need to find the largest one is this sub-tree
	if root.Left != nil {
		return FindLargest(root)
	}

	// If there isn't a left or right, then we are currently at the largest node
	// as we are traversing to the right (larger) node each time.
	// Because of this, the node we visited before must be the second largest
	return previousNodeValue
}

func FindLargest(root *Node) int {

	if root.Right == nil {
		return root.value
	}

	return FindLargest(root.Right)
}
