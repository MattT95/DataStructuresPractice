package main

import (
	"fmt"
)

/*
	Write a method to replace all the spaces in a string with "%20".
	You may assume that the string has sufficient space at the end
	to hold the additional characters, and that you are given the "true"
	length of the string
*/

func main() {
	example1 := []rune("Mr John Smith    ")
	example2 := []rune(" Mr John Smith      ")
	example3 := []rune("Mr John Smith       ")

	fmt.Println(URLify(example1, 13))
	fmt.Println(URLify(example2, 14))
	fmt.Println(URLify(example3, 14))
}

/*
	We can iterate through the string from left to right,
	and replace each space we see along the way.

	On each occurrence of a space character, we can shift the "tail" contents
	of the string (the section we have not yet iterated over) to the right by 2,
	and from this we can make space to insert the url-encoded spaces.

	By doing this on the array "in-place", we can save on the string copying that
	would occur if we were to concatenate substrings of the "tail", encoding,
	and head of the array. These copies would require us iterating over the string
	multiple times, and lead to worst-case quadratic runtime complexity, in proportion to the size
	of the string (as we are iterating over the string during a copy as a side-effect from concatenation)
*/
func URLify(str []rune, length int) string {

	idx := 0
	for idx < len(str) {
		nextCharacter := str[idx]

		if nextCharacter == ' ' {
			for j := len(str) - 1; j > idx+2; j-- {
				str[j] = str[j-2]
			}

			str[idx] = '%'
			str[idx+1] = '2'
			str[idx+2] = '0'
			idx = idx + 2
		}

		idx++
	}

	return string(str)
}
