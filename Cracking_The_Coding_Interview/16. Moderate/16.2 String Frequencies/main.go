package main

import "fmt"

/*
	Word Frequencies:
		Design a method to find the frequency of occurrences of any given word in a book.
		What if we were running this algorithm multiple times?
*/

func main() {

	book := "This is a test string with some duplicate test and another test and another"

	// Expected occurrences of "test" is 3
	fmt.Println(CountOccurrences("test", book))

	// Expected occurrences of "with" is 1
	fmt.Println(CountOccurrences("with", book))

	// Expected occurrences of "and" is 2
	fmt.Println(CountOccurrences("and", book))
}

func CountOccurrences(word string, book string) int {

	// For single use, an iteration based on analysing each word in the string is sufficient
	// This is O(n) where n is the length of the string
	// (not the words, as we need to analyse each character in the string/array to determine splitting)

	count := 0
	currentPos := 0

	for {
		if currentPos > len(book)-1 {
			break
		}

		if book[currentPos] != byte(' ') {
			nextWordEndPos := FindEndOfNextWord(currentPos, book)

			if WordMatches(word, book, currentPos, nextWordEndPos) {
				count++
			}
			currentPos = nextWordEndPos + 1
		} else {
			currentPos++
		}

	}

	return count
}

func FindEndOfNextWord(startPos int, str string) int {

	delemiterChar := byte(' ')
	currentPos := startPos

	for {
		if currentPos == len(str)-1 || str[currentPos+1] == delemiterChar {
			return currentPos
		}

		currentPos++
	}
}

func WordMatches(word string, str string, substringStartIdx int, substringEndIdx int) bool {

	if len(word) != (substringEndIdx - substringStartIdx) {
		return false
	}

	for i := 0; substringStartIdx+i <= substringEndIdx; i++ {
		fmt.Printf("%s \n", rune(str[substringStartIdx+i]))

		if word[i] != str[substringStartIdx+i] {
			return false
		}
	}

	return true
}
