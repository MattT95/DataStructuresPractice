package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	tree := CreateBinaryTree(array, 0, len(array)-1)

	tree.Print()
}

// Worst Case Tree:
//	1
// 		2
// 			3
// 				4
// 					5

// Best Case Tree:
// 					3
// 				2		4
// 			1				5
//
//
//

func CreateBinaryTree(input []int, start int, end int) Node {

	middle := start + (end-start)/2

	tree := Node{input[middle], nil, nil}

	if start != middle {
		node := CreateBinaryTree(input, start, middle-1)
		tree.Left = &node
	}

	if end != middle {
		node := CreateBinaryTree(input, middle+1, end)
		tree.Right = &node
	}

	return tree
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) Print() {
	queue := make([]*Node, 0)
	queue = append(queue, node)

	for {
		if len(queue) == 0 {
			break
		}

		next := queue[0]
		queue = queue[1:]
		fmt.Printf("Root: (%p) %d, Left: %p, Right: %p \n", next, (*next).Value, (*next).Left, (*next).Right)

		if next.Left != nil {
			queue = append(queue, next.Left)
		}

		if next.Right != nil {
			queue = append(queue, next.Right)
		}
	}
}
