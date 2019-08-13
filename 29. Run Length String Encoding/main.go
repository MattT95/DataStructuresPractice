package main

import (
	"fmt"
	"strconv"
)

// Run-length encoding is a fast and simple method of encoding strings.
// The basic idea is to represent repeated successive characters as a single count and character.
// For example, the string "AAAABBBCCDAA" would be encoded as "4A3B2C1D2A".
//
// Implement run-length encoding and decoding.
// You can assume the string to be encoded have no digits
// and consists solely of alphabetic characters.
// You can assume the string to be decoded is valid.

// Linear time for encode + decode

func main() {
	fmt.Println("statting")

	example := "AAAABBBCCDAA"

	encoded := Encode(example)
	fmt.Println(encoded)

	decoded := Decode(encoded)
	fmt.Println(decoded)
}

func Encode(str string) string {

	if len(str) == 0 {
		return ""
	}

	encodedString := ""

	previousChar := str[0]
	previousCharOccurrences := 1

	for i := 1; i < len(str); i++ {
		currentChar := str[i]

		if currentChar != previousChar {

			encodedString += strconv.Itoa(previousCharOccurrences)
			encodedString += string(previousChar)
			previousChar = currentChar
			previousCharOccurrences = 1
			continue
		}

		previousCharOccurrences++
	}

	encodedString += strconv.Itoa(previousCharOccurrences)
	encodedString += string(previousChar)

	return encodedString
}

func Decode(str string) string {

	decodedString := ""

	if len(str) == 0 {
		return decodedString
	}

	for i := 0; i < len(str); i = i + 2 {

		occurrences, _ := strconv.Atoi(string(str[i]))
		character := str[i+1]

		for j := 0; j < occurrences; j++ {
			decodedString += string(character)
		}
	}

	return decodedString
}
