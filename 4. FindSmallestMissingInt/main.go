package main

import (
	"fmt"
)

func main() {

	tests := [][]int { 
		{ 0, 1, 2, -1 },
		{ 1, 2, 4, 3},
		{ 1, 2, 2, 4},
		{ },
		{ 0, 0, 0, 0},
		{ 1, -1, 1, -1 },
		{ 0, 1, -1, 2, 3, 4, 5, 7, 9 },
	}

	for _, val := range tests {
		smallestInt := findSmallestMissingInt(val)
		fmt.Println(smallestInt) 
	}
}

func findSmallestMissingInt(list []int) int{
	
	smallestInt := 1

	for {
		foundZero := false

		for idx, _ := range list {
			list[idx] = list[idx] - 1
			
			if list[idx] == 0 {
				foundZero = true
			}
		}

		if foundZero == true {
			smallestInt++
		} else {
			break;
		}
	}

	return smallestInt
}
