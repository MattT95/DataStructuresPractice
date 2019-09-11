package main

import (
	"fmt"
)

/*

   a
  /|\
 b c d
    / \
   e   f
  / \
 g   h

a-b: 3, a-c: 5, a-d: 8, d-e: 2, d-f: 4, e-g: 1, e-h: 1,
the longest path would be c -> a -> d -> f,
with a length of 17.

*/

func main() {

	g := Node{1, []*Node{}}
	h := Node{1, []*Node{}}
	e := Node{2, []*Node{&g, &h}}
	f := Node{4, []*Node{}}
	d := Node{8, []*Node{&e, &f}}
	c := Node{5, []*Node{}}
	b := Node{3, []*Node{}}
	a := Node{0, []*Node{&b, &c, &d}}

	longest := FindLongestPath(&a, 0)

	fmt.Println(longest)
}

func FindLongestPath(root *Node, currentPathSize int) int {

	var firstBiggestNode *Node
	var firstBiggestVal int
	var secondBiggestNode *Node
	var secondBiggestVal int

	for i := 0; i < len(root.Nodes); i++ {
		node := ((*root).Nodes[i])

		currentVal := FindLongestPath(node, currentPathSize+root.Value)

		if currentVal > firstBiggestVal {

		}
	}
}

type Node struct {
	Value int
	Nodes []*Node
}
