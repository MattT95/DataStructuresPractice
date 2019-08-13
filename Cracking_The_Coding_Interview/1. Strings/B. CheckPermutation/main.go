package main

import (
	"fmt"
)

/*
	Given two strings, write a method to decide if one if a
	permutation of the other
*/

func main() {

	input := "abcdef"
	permutation := "fedcba"
	nonPermutation := "abcdfg"

	fmt.Println(CheckPermutation(input, permutation))
	fmt.Println(CheckPermutation(input, nonPermutation))
}

func CheckPermutation(input string, compare string) bool {

	if len(input) != len(compare) {
		return false
	}

	occurences := make(map[rune]int)

	for i := 0; i < len(input); i++ {
		character := rune(input[i])
		_, hasEntry := occurences[character]

		if hasEntry {
			occurences[character] = occurences[character] + 1
		} else {
			occurences[character] = 1
		}
	}

	for i := 0; i < len(compare); i++ {
		character := rune(compare[i])
		_, hasEntry := occurences[character]

		if !hasEntry {
			return false
		}

		if occurences[character] == 1 {
			delete(occurences, character)
		} else {
			occurences[character] = occurences[character] - 1
		}
	}
	return true
}
