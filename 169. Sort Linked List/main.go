package main

import "fmt"

/*
Given a linked list, sort it in O(n log n) time and constant space.

For example, the linked list
	4 -> 1 -> -3 -> 99
should become
	-3 -> 1 -> 4 -> 99.
*/

func main() {
	list := LinkedList{nil}

	list.Add(4)
	list.Add(1)
	list.Add(-3)
	list.Add(99)

	list.Print()

	SortList(&list)

	list.Print()
}

func SortList(list *LinkedList) {

	firstNode := SortListSegment((*list).Head, nil)

	(*list).Head = firstNode
}

func SortListSegment(startNode *Node, endNode *Node) *Node {

	fmt.Printf("Running sort %d %d", startNode, endNode)
	if startNode == endNode {
		return startNode
	}

	pivot := startNode
	current := (*pivot).Next
	lastVisited := pivot
	firstLeftNode := pivot
	firstRightNode := pivot
	lastLeftNode := pivot
	lastRightNode := pivot

	for {
		if current == endNode {
			break
		}

		fmt.Println((*current).Value)
		if (*pivot).Value > (*current).Value {
			(*lastVisited).Next = (*current).Next
			(*current).Next = firstLeftNode
			firstLeftNode = current
			if lastLeftNode == pivot {
				lastLeftNode = current
			}
		} else {
			if firstRightNode == pivot {
				firstRightNode = current
			}
			lastRightNode = current
			lastVisited = current
		}

		current = (*lastVisited).Next
	}
	// fmt.Println("POS")
	// fmt.Println(firstLeftNode)

	// fmt.Println(lastLeftNode)

	// fmt.Println(firstRightNode)

	// fmt.Println(lastRightNode)
	// return firstLeftNode

	fmt.Println("sorting left")
	sortedLeftHead := SortListSegment(firstLeftNode, lastLeftNode)
	fmt.Println("sorting right")
	sortedRightHead := SortListSegment(firstRightNode, lastRightNode)

	if sortedRightHead != pivot {
		pivot.Next = sortedRightHead
	}

	return sortedLeftHead
}

type LinkedList struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (list *LinkedList) Add(val int) {

	node := Node{val, nil}

	if (*list).Head == nil {
		(*list).Head = &node
	} else {

		current := (*list).Head

		for {
			if (*current).Next == nil {
				(*current).Next = &node
				break
			}
			current = (*current).Next
		}
	}
}

func (list *LinkedList) Print() {

	current := (*list).Head

	for {
		if current == nil {
			break
		}

		fmt.Printf("%d ", (*current).Value)

		current = (*current).Next
	}

	fmt.Println()
}
