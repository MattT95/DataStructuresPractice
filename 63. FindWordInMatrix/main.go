package main

import (
	"fmt"
)

func main() {

	word := "AB"

	matrix := [][]rune{
		{'F', 'A', 'C', 'I'},
		{'O', 'B', 'Q', 'P'},
		{'A', 'N', 'O', 'B'},
		{'M', 'A', 'S', 'S'},
	}

	wordExists := DoesWordExist(word, matrix, 0, 0)

	fmt.Println(wordExists)
}

func DoesWordExist(word string, wordMatrix [][]rune, checkPosX int, checkPosY int) bool {

	for i := 0; i < len(wordMatrix); i++ {
		if DoesWordExistInRow(i, word, wordMatrix) || DoesWordExistInColumn(i, word, wordMatrix) {
			return true
		}
	}
	return false
}

func DoesWordExistInRow(row int, word string, wordMatrix [][]rune) bool {

	if len(word) >= len(wordMatrix[row]) {
		return false
	}

	for j := 0; j < len(word); j++ {
		if rune(word[j]) != wordMatrix[row][j] {
			return false
		}
	}
	return true
}

func DoesWordExistInColumn(column int, word string, wordMatrix [][]rune) bool {

	if len(word) >= len(wordMatrix) {
		return false
	}

	for j := 0; j < len(word); j++ {
		if column >= len(wordMatrix[j]) {
			return false
		}

		if rune(word[j]) != wordMatrix[j][column] {
			return false
		}
	}
	return true
}
