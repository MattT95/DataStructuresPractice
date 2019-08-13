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
	for i := 0; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				tmp := array[j-1]
				array[j-1] = array[j]
				array[j] = tmp
			}
		}
	}

	return array
}
