package main

import (
	"fmt"
)

func main() {

	array := make([]int, 0)

	// array := []int{3, 2, 1}
	// fmt.Println(Sort(array))

	array = []int{6, 7, 2, 1, 9, 20, 25}
	fmt.Println(Sort(array))
}

func Sort(array []int) []int {

	// Convert the list to a max heap
	for i := (len(array) / 2) - 1; i >= 0; i-- {
		fmt.Println(i)
		Heapify(&array, i)

	}

	result := make([]int, len(array))

	for i := 0; i < len(result); i++ {
		result[i] = array[0]
		array[0] = array[len(array)-(i+1)]
		array[len(array)-(i+1)] = -1
		Heapify(&array, 0)
	}

	return result
}

func Heapify(array *[]int, rootIdx int) {
	fmt.Println(*array)
	leftIdx := (rootIdx * 2) + 1
	rightIdx := (rootIdx * 2) + 2
	largestIdx := rootIdx

	if leftIdx < len(*array) && (*array)[leftIdx] > (*array)[largestIdx] {
		largestIdx = leftIdx
	}

	if rightIdx < len(*array) && (*array)[rightIdx] > (*array)[largestIdx] {
		largestIdx = rightIdx
	}

	if largestIdx != rootIdx {
		temp := (*array)[rootIdx]
		(*array)[rootIdx] = (*array)[largestIdx]
		(*array)[largestIdx] = temp

		Heapify(array, largestIdx)
	}
}
