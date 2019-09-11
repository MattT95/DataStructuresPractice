package main

import (
	"fmt"
)

// You are given an array of length n + 1 whose elements belong to the set
// {1, 2, ..., n}. By the pigeonhole principle, there must be a duplicate.
// Find it in linear time and space.

func main() {

	test := []int{1, 2, 4, 3, 4, 5}
	fmt.Println(GetDuplicate(test))
}

func GetDuplicate(input []int) int {

	counts := make([]int, len(input))

	for i := 0; i < len(input); i++ {

		current := input[i]
		currentCount := counts[current]

		if currentCount == 1 {
			return current
		}

		counts[current] = counts[current] + 1
	}

	return -1
}
