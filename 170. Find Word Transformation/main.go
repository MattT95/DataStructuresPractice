package main

import (
	"container/list"
	"fmt"
)

// Given a start word, an end word, and a dictionary of valid words,
// find the shortest transformation sequence from start to end such that only one letter is changed at each step of the sequence, and each transformed word exists in the dictionary. If there is no possible transformation, return null.
// Each word in the dictionary have the same length as start and end and is lowercase.

// For example, given start = "dog", end = "cat", and dictionary = {"dot", "dop", "dat", "cat"},
// return ["dog", "dot", "dat", "cat"].

// Given start = "dog", end = "cat",
// and dictionary = {"dot", "tod", "dat", "dar"},
// return null as there is no possible transformation from dog to cat.

func main() {
	start := "dog"
	end := "cat"
	dictionary := map[string]bool{
		"dog": false,
		"dot": false,
		"dat": false,
		"cat": false,
	}

	path := FindPath(start, end, dictionary)

	if path.Front() == nil {
		fmt.Println("Emptys")
	}

	fmt.Println(path.Len())
	for i := (*path).Front(); i != nil; i = i.Next() {
		fmt.Printf("%s ", (*i).Value.(string))
	}
}

func FindPath(start string, end string, dict map[string]bool) *list.List {
	queue := make([]Node, 0)

	startNode := Node{start, list.New()}
	(*startNode.List).PushBack(start)

	queue = append(queue, startNode)
	words := make([]string, len(dict))

	i := 0
	for k := range dict {
		words[i] = k
		i++
	}

	for {
		if len(queue) == 0 {
			break
		}

		current := queue[0]

		if current.Value == end {
			return current.List
		}

		queue = queue[1:]

		for i := 0; i < len(words); i++ {
			if HasOneDifference(current.Value, words[i]) {
				val, ok := dict[words[i]]

				if ok && val != true {
					dict[words[i]] = true
					copiedList := CopyList(current.List)
					(*copiedList).PushBack(words[i])

					queue = append(queue, Node{words[i], copiedList})
				}
			}
		}

	}

	fmt.Println("out")
	return list.New()
}

func CopyList(in *list.List) *list.List {
	out := list.New()

	for i := (*in).Front(); i != nil; i = i.Next() {
		(*out).PushBack((*i).Value)
	}

	return out
}

func HasOneDifference(strA string, strB string) bool {
	diff := 0
	for i := 0; i < len(strA); i++ {
		if strA[i] != strB[i] {
			if diff == 1 {
				return false
			}
			diff++
		}
	}
	return diff == 1
}

type Node struct {
	Value string
	List  *list.List
}
