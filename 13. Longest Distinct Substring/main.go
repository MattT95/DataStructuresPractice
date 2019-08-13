package main

import "fmt"

// This problem was asked by Amazon.
// Given an integer k and a string s,
// find the length of the longest substring that contains at most k distinct characters.

// For example,
// given s = "abcba" and k = 2,
// the longest substring with k distinct characters is "bcb".

func main() {

	s := "abcccbbbbba"
	k := 3

	fmt.Println(FindLongestSubstring(s, k))
	fmt.Println(FindLongestSubstringFaster(s, k))
}

func FindLongestSubstringFaster(str string, k int) int {

	maxLength := 0
	longestString := ""
	distinctCharacters := make(map[byte]int)

	for i := 0; i < len(str); i++ {

		currentChar := str[i]

		for {
			_, alreadyHasCharacter := distinctCharacters[currentChar]

			if alreadyHasCharacter || len(distinctCharacters) < k {
				break
			}

			headChar := longestString[0]

			if distinctCharacters[headChar] == 1 {
				delete(distinctCharacters, byte(headChar))
			} else {
				distinctCharacters[headChar] = distinctCharacters[headChar] - 1
			}

			if len(longestString) > 0 {
				longestString = longestString[1:]
			}
		}

		distinctCharacters[currentChar] = distinctCharacters[currentChar] + 1
		longestString = longestString + string(currentChar)
		if len(longestString) > maxLength {
			maxLength = len(longestString)
		}
	}
	return maxLength

}

func FindLongestSubstring(str string, k int) int {

	max := 0

	for i := 0; i < len(str); i++ {
		count := 0
		distinctCharacters := make([]byte, len(str))
		distinctCount := 0

		for j := i; j < len(str); j++ {

			currentChar := str[j]
			isNewDistinctChar := true

			for _, distinctChar := range distinctCharacters {
				if distinctChar == currentChar {
					isNewDistinctChar = false
					break
				}
			}

			if isNewDistinctChar {
				distinctCharacters[j] = currentChar
				distinctCount++
			}

			if distinctCount > k {
				break
			}
			count++
		}

		if count > max {
			max = count
		}
	}

	return max
}
