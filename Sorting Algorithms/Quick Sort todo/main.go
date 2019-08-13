package main

import (
	"fmt"
)

func main() {

	array := []int{3, 2, 1}
	Sort(&array)
	fmt.Println(array)

	array = []int{6, 7, 2, 1, 20, 9, 3}
	Sort(&array)
	fmt.Println(array)
}

func Sort(array *[]int) {
	SortArraySection(array, 0, len(*array)-1)
}

func SortArraySection(array *[]int, startIdx int, endIdx int) {

	if startIdx < endIdx && endIdx > startIdx {
		endPivotIdx := PartitionBasedOnPivot(array, startIdx, endIdx, startIdx)

		fmt.Printf("Start: %d, End: %d, Pivot: %d  \n", startIdx, endIdx, endPivotIdx)
		SortArraySection(array, startIdx, endPivotIdx-1)
		SortArraySection(array, endPivotIdx+1, endIdx)
	}
}

func PartitionBasedOnPivot(array *[]int, start int, end int, pivotIdx int) int {

	pivot := (*array)[pivotIdx]

	for i := start + 1; i <= end; i++ {
		if pivot > (*array)[i] {
			temp := (*array)[i]
			ShiftRight(array, pivotIdx, i)
			(*array)[pivotIdx] = temp
			pivotIdx++
			fmt.Println("Move value left")
		} else {
			fmt.Println("Value stays")
		}
		fmt.Println(*array)
	}

	return pivotIdx
}

func ShiftRight(array *[]int, start int, end int) {
	for i := end; i > start; i-- {
		(*array)[i] = (*array)[i-1]
	}
}
