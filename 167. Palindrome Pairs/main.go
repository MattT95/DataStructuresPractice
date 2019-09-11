package main

// This problem was asked by Airbnb.
// Given a list of words, find all pairs of unique indices
// such that the concatenation of the two words is a palindrome.

// For example, given the list ["code", "edoc", "da", "d"], return [(0, 1), (1, 0), (2, 3)].
import (
	"container/list"
	"fmt"
)

func main() {
	// ["code", "edoc", "da", "d"]
	array := []string{"code", "edoc", "da", "d"}

	pairs := *PalindromePairs(array)

	for node := pairs.Front(); node != nil; node = node.Next() {
		fmt.Printf("(%d, %d) \n", ((*node).Value.([]int))[0], ((*node).Value.([]int))[1])
	}
}

func PalindromePairs(array []string) *list.List {
	pairs := list.New()

	for i := 0; i < len(array)-1; i++ {
		for j := i; j < len(array); j++ {

			stringA := array[i]
			stringB := array[j]

			if IsPalindrome(stringA + stringB) {
				pairs.PushBack([]int{i, j})
			}

			if IsPalindrome(stringB + stringA) {
				pairs.PushBack([]int{j, i})
			}
		}
	}

	return pairs
}

func IsPalindrome(str string) bool {

	i := 0
	j := len(str) - 1

	for {
		if i == j || i > j {
			return true
		}

		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}
}
