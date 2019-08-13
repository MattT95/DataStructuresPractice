package main

import (
	"fmt"
)

func main() {

	array := []int{3, 2, 1}
	fmt.Println(Sort(array))

	array = []int{6, 7, 2, 1, 9, 20}
	fmt.Println(Sort(array))
}

func Sort(array []int) []int {

	if len(array) == 1 {
		return array
	}

	leftArray := make([]int, len(array)/2)
	for i := 0; i < len(leftArray); i++ {
		leftArray[i] = array[i]
	}

	rightArray := make([]int, len(array)/2+len(array)%2)
	for i := 0; i < len(rightArray); i++ {
		rightArray[i] = array[len(array)/2+i]
	}

	sortedLeftArray := Sort(leftArray)
	sortedRightArray := Sort(rightArray)

	return Merge(sortedLeftArray, sortedRightArray)
}

func Merge(leftArray []int, rightArray []int) []int {

	result := make([]int, len(leftArray)+len(rightArray))

	leftIdx := 0
	rightIdx := 0
	for i := 0; i < len(result); i++ {

		if rightIdx >= len(rightArray) {
			result[i] = leftArray[leftIdx]
			leftIdx++
		} else if leftIdx >= len(leftArray) {
			result[i] = rightArray[rightIdx]
			rightIdx++
		} else if leftArray[leftIdx] < rightArray[rightIdx] {
			result[i] = leftArray[leftIdx]
			leftIdx++
		} else {
			result[i] = rightArray[rightIdx]
			rightIdx++
		}
	}

	return result
}
