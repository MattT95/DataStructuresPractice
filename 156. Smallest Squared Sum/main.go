package main

import "fmt"

// Given a positive integer n, find the smallest number of squared integers which sum to n.

// For example, given n = 13, return 2 since 13 = 32 + 22 = 9 + 4.

// Given n = 27, return 3 since 27 = 32 + 32 + 32 = 9 + 9 + 9.

func main() {

	fmt.Println(SmallestSquareSum(13))

	fmt.Println(SmallestSquareSum(1350))
}

func SmallestSquareSum(sum int) int {

	queue := Queue{nil, nil}

	largestSquareRoot := FindLargestSquareRoot(sum)
	fmt.Printf("Largest Square Root: %d \n", largestSquareRoot)

	for i := largestSquareRoot; i >= 1; i-- {
		queue.Add(sum-(i*i), 1)
	}

	for {
		if queue.Empty() {
			return -1
		}

		val, count := queue.Pop()

		fmt.Printf("[Queue] %d - %d \n", val, count)
		if val == 0 {
			fmt.Println("DONE")
			return count
		}

		if val < 0 {
			continue
		}

		largestSquareRoot = FindLargestSquareRoot(val)

		for i := largestSquareRoot; i >= 1; i-- {
			queue.Add(val-(i*i), count+1)
		}
	}
}

func FindLargestSquareRoot(num int) int {

	start := 1
	end := num

	for {
		current := start + ((end - start) / 2)
		fmt.Printf("[Root]: %d \n", current)
		currentSquare := current * current
		nextCurrentSquare := (current + 1) * (current + 1)

		if currentSquare <= num && nextCurrentSquare >= num {
			return current
		}

		if currentSquare >= num {
			end = current - 1
		} else {
			start = current + 1
		}
	}
}

type Queue struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value int
	Count int
	Next  *Node
}

func (queue *Queue) Add(val int, count int) {
	node := Node{val, count, nil}

	if queue.Head == nil {
		queue.Head = &node
		queue.Tail = &node
	} else {
		queue.Tail.Next = &node
		queue.Tail = &node
	}
}

func (queue *Queue) Pop() (int, int) {

	val := (*(*queue).Head)

	(*queue).Head = val.Next

	return val.Value, val.Count
}

func (queue *Queue) Empty() bool {
	return (*queue).Head == nil
}
