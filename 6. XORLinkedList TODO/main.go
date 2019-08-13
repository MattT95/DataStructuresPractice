package main

import (
	"strconv"
	"unsafe"
	"fmt"
)

func main() {
	
	list := LinkedList{}
	
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)

	list.Print()

	// fmt.Println(list.Get(2)) // 3
	// fmt.Println(list.Get(5)) // 6
	// fmt.Println(list.Get(6)) // -1
}

type LinkedList struct{
	Head *Node
	Tail *Node
}

type Node struct {
	value int
	both *uintptr
}

func (list *LinkedList) Add(value int) {
	
	if list.Head == nil {
		node := Node{ value, nil }
		list.Head = &node
		list.Tail = &node	
		return
	} 
	
	tailPtr := uintptr(unsafe.Pointer(list.Tail))	
	newNode := Node{ value, &tailPtr }
	newNodePtr  := uintptr(unsafe.Pointer(&newNode))
		
	if list.Tail.both == nil {
		list.Tail.both = &newNodePtr
	} else {
		newBothPtr := *(list.Tail.both) | newNodePtr 
		list.Tail.both = &newBothPtr
	} 

	list.Tail = &newNode
}

func (list *LinkedList) Get(position int) int{
	
	count := 0
	var previousNode *Node = nil
	currentNode := list.Head

	for {
		if currentNode == nil {
			return -1
		}

		if count == position {
			return currentNode.value
		}

		currentNodePtr := uintptr(unsafe.Pointer(&currentNode))

		var nextNodePtr uintptr
		if previousNode == nil {
			nextNodePtr = *currentNode.both
		} else {
			previousNodePtr := uintptr(unsafe.Pointer(&previousNode))
			nextNodePtr = currentNodePtr | previousNodePtr
		}

		currentNode = (*Node)(unsafe.Pointer(nextNodePtr))
		previousNode = currentNode
		
		count++
	}
}

func (list *LinkedList) Print() {

	var previousNode *Node = nil
	currentNode := list.Head
	
	str := ""

	for {
		fmt.Println(str)
		fmt.Println("--")
		if currentNode == nil {
			break
		}

		str += " " + strconv.Itoa(currentNode.value)

		if previousNode == nil {
			previousNode = currentNode
			currentNode = (*Node)(unsafe.Pointer(currentNode.both))
			continue
		}

		previousNodePtr := uintptr(unsafe.Pointer(&previousNode))
		currentNodeBothPtr := *currentNode.both

		fmt.Println("<-->")
		fmt.Println("previous")
		fmt.Println(previousNodePtr)
		fmt.Println("current both")
		fmt.Println(currentNodeBothPtr)
		fmt.Println("<-->")
		if previousNodePtr == currentNodeBothPtr {
			currentNode = nil
		} else {
			currentNodePtr := previousNodePtr | currentNodeBothPtr
			previousNode = currentNode
			currentNode = (*Node)(unsafe.Pointer(currentNodePtr))
		}
		
	}

	fmt.Println(str)
}