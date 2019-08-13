package main

import (
	"fmt"
)

// We can determine how "out of order" an array A is by counting the number of inversions
// it has. Two elements A[i] and A[j] form an inversion
// if A[i] > A[j] but i < j. That is, a smaller element appears after a larger element.

// Given an array, count the number of inversions it has. Do this faster than O(N^2) time.

// You may assume each element in the array is distinct.
// For example, a sorted list has zero inversions.

// The array [2, 4, 1, 3, 5] has three inversions:
// (2, 1), (4, 1), and (4, 3).

// The array [5, 4, 3, 2, 1] has ten inversions: every distinct pair forms an inversion.

// 0 1 2 3 4
// ---------------
// 5 4 3 2 1

// 4 3 2 1 0

// 8 7 6 5 4 2 3 1 0
 
// 0 1 3 2 - - - - - 

// 1 + 2 + 2 + 
// [,2,3,4,5,6,7]

// [3,1] => (3, 1)
// [2,3,1] => (2,1)
// [4,2,3,1] => (4,2), (4,3), (4,1) 


func main() {

	arrayA := []int{2, 4, -1, 3, 5}
	fmt.Println(CountInversions(arrayA))

	arrayB := []int{5, 4, 3, 2, 1}
	fmt.Println(CountInversions(arrayB))
}

func CountInversionsBruteForce(array []int) int {

	count := 0

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				count++
			}
		}
	}

	return count
}

func CountInversions(array []int) int {

}
