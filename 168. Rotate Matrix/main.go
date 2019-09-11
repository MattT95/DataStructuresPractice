package main

import "fmt"

// Given an N by N matrix, rotate it by 90 degrees clockwise.

// For example, given the following matrix:

// [[1, 2, 3],
//  [4, 5, 6],
//  [7, 8, 9]]
// you should return:

// [[7, 4, 1],
//  [8, 5, 2],
//  [9, 6, 3]]
// Follow-up: What if you couldn't use any extra space?

func main() {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	result := Rotate90(input)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result); j++ {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Printf("\n")
	}
}

func Rotate90(input [][]int) [][]int {

	for i := 0; i < len(input)/2; i++ {

		start := i
		end := len(input) - (i + 1)
		for j := start; j < end; j++ {

			top := input[i][j]
			right := input[j][len(input)-(1+i)]
			bottom := input[len(input)-(1+i)][len(input)-(1+j)]
			left := input[len(input)-(1+j)][i]

			input[i][j] = left
			input[j][len(input)-(1+i)] = top
			input[len(input)-(1+i)][len(input)-(1+j)] = right
			input[len(input)-(1+j)][i] = bottom
		}
	}

	return input
}
