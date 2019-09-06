package main

import "fmt"

/*
* Gray code is a binary code where each successive value differ in only one bit, as well as when wrapping around.
* Gray code is common in hardware so that we don't see temporary spurious values during transitions.
*
* Given a number of bits n, generate a possible gray code for it.
*
* For example, for n=2 one gray code would be:
*
* [00, 01, 11, 10]
*
 */

// Generate the numbers and iterate in LIFO order to generate next with additional bit

func main() {
	bits := 4

	grayCode := Stack{}

	GenerateGrayCode(&grayCode, bits)

	grayCode.Print()
}

func GenerateGrayCode(code *Stack, bits int) {

	if bits == 1 {
		(*code).Push(0)
		(*code).Push(1)
		return
	}

	GenerateGrayCode(code, bits-1)

	currentNode := (*code).Head

	for {
		if currentNode == nil {
			break
		}

		code.Push((*currentNode).Value + (1 << uint(bits-1)))

		currentNode = (*currentNode).Next
	}
}

type Stack struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (stack *Stack) Push(val int) {
	node := Node{val, nil}

	if stack.Head == nil {
		(*stack).Head = &node
	} else {
		temp := (*stack).Head
		(*stack).Head = &node
		(*(*stack).Head).Next = temp
	}
}

func (stack *Stack) Print() {
	currentNode := stack.Head

	for {
		if currentNode == nil {
			break
		}

		fmt.Printf("%b\n", (*currentNode).Value)

		currentNode = (*currentNode).Next
	}
}
