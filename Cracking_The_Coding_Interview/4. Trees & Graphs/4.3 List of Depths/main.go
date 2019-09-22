package main

import "container/list"

// List of Depths
// 		Given a binary tree, design an algorithm which creates a linked list of all the nodes at each depth
// 		(e.g., if you have a tree with depth D, you'll have D linked lists).

//
// Example:
//		For tree
//
//						1
//				2				3
//			4		5		6		7
//
// 		The output should be
//			1
//			2 -> 3
//			4 -> 5 -> 6 -> 7

func main() {

	g := Node{7, nil, nil}
	f := Node{6, nil, nil}
	e := Node{5, nil, nil}
	d := Node{4, nil, nil}

	c := Node{3, &f, &g}
	b := Node{2, &d, &e}

	root := Node{1, &b, &c}

	lists := GetListOfDepths(&root)
}

func GetListOfDepths(list *List, root *Node) list.List {
	queue := Queue{}

	result := list.New()

	queue.Add(QueueWorkItem{0, &root})

	for {
		next := queue.Take()

		if next == nil {
			break
		}

		(*result).Get(depth).Add((*next).Value)

		queue.Push(QueueWorkItem{(*next).Depth + 1, (*next).Left})
		queue.Push(QueueWorkItem{(*next).Depth + 1, (*next).Right})
	}

	return result
}

type Queue struct {
	list *LinkedList
}

type QueueWorkItem struct {
	Depth int
	Node  *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type LinkedList struct {
	Head *LinkedListNode
	Last *LinkedListNode
}

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

func (linkedList *LinkedList) Add(value int) {
	node := LinkedListNode{value, nil}

	if (*linkedList).Head == nil {
		linkedList.Head = &node
		linkedList.Last = &node
	} else {
		linkedList.Last.Next = &node
		linkedList.Last = &node
	}
}
