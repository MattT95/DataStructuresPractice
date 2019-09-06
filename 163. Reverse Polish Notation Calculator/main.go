package main

import (
	"fmt"
)

// The expression is given as a list of numbers and operands.
// For example: [5, 3, '+'] should return 5 + 3 = 8.
// For example, [15, 7, 1, 1, '+', '-', '/', 3, '*', 2, 1, 1, '+', '+', '-'] should return 5, since it is equivalent to ((15 / (7 - (1 + 1))) * 3) - (2 + (1 + 1)) = 5.

func main() {
	test := []interface{}{5, 3, '+'}

	fmt.Println(Calculate(test))

	test = []interface{}{15, 7, 1, 1, '+', '-', '/', 3, '*', 2, 1, 1, '+', '+', '-'}

	fmt.Println(Calculate(test))
}

func Calculate(input []interface{}) int {

	stack := Stack{nil}

	for i := 0; i < len(input); i++ {

		current := input[i]

		number, isInt := current.(int)

		if isInt {
			stack.Push(number)
		} else {
			operator := current.(rune)
			right := stack.Pop()
			left := stack.Pop()
			num := RunArithmetic(left, right, operator)
			stack.Push(num)
		}
	}

	return stack.Pop()
}

func RunArithmetic(leftOperand int, rightOperand int, operator rune) int {
	switch operator {
	case '*':
		return leftOperand * rightOperand
	case '/':
		return leftOperand / rightOperand
	case '+':
		return leftOperand + rightOperand
	case '-':
		return leftOperand - rightOperand
	}

	panic("Invalid operator")
}

type Stack struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (stack *Stack) Push(val int) {

	if stack.Head == nil {
		stack.Head = &Node{val, nil}
	} else {
		temp := stack.Head
		stack.Head = &Node{val, temp}
	}
}

func (stack *Stack) Pop() int {

	if stack.Head == nil {
		panic("no node to pop.")
	}

	result := *(stack.Head)

	stack.Head = result.Next

	return result.Value
}
