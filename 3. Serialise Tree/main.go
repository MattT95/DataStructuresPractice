package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println("Running tests")

	example := Node{
		"root",
		&Node { 
			"left", 
			&Node { "left.left", nil, nil },
			nil,
		},	
		&Node { "right", nil, nil },
	}

	serialisedTree := Serialise(&example)
	
	fmt.Println(serialisedTree)

	deserialisedTree := Deserialise(serialisedTree)

	fmt.Println(deserialisedTree.Value)
	fmt.Println(deserialisedTree.Left.Value)
	fmt.Println(deserialisedTree.Left.Right)
	fmt.Println(deserialisedTree.Left.Left.Value)
	fmt.Println(deserialisedTree.Left.Left.Left)
	fmt.Println(deserialisedTree.Left.Left.Right)
	fmt.Println(deserialisedTree.Right.Value)
}

func Serialise(node *Node) string {

	if node == nil {
		return "#"
	}

	return  (*node).Value + "-" + Serialise((*node).Left) + "-" + Serialise((*node).Right) 
}

func Deserialise(input string) Node {

	elements := strings.Split(input, "-")

	queue := Queue{}

	for _, element := range elements {
		queue.Push(element)
	}

	return *DeserialiseElements(&queue)
}

func DeserialiseElements(queue *Queue) *Node {
	
	next := queue.Pop()

	if *next == "#" {
		return nil
	}
	
	return &Node{
		*next,
		DeserialiseElements(queue),
		DeserialiseElements(queue)} 
}

type Queue struct {
	current *QueueElement
	last *QueueElement
}

func (queue *Queue) Push(value string) {
	if queue.current == nil {
		queue.current = &QueueElement{ value, nil }
		queue.last = queue.current
	} else {
		newLast := QueueElement{ value, nil }
		queue.last.Next = &newLast
		queue.last = &newLast
	}
}

func (queue *Queue) Pop() *string {
	if queue.current == nil {
		return nil
	}
	
	value := queue.current.Value
	queue.current = queue.current.Next
	
	return &value
}

type QueueElement struct {
	Value string
	Next *QueueElement
}

type Node struct {
	Value string
	Left *Node
	Right *Node
}