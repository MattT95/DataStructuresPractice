package main

import (
	"fmt"
	"strings"
)

// Implement an autocomplete system.
// That is, given a query string s and a set of all possible query strings,
// return all strings in the set that have s as a prefix.

// For example, given the query string de and the set of strings [dog, deer, deal],
// return [deer, deal].

func main() {

	word := "de"
	dictionary := []string{"dog", "deer", "deal"}

	result := AutoComplete(word, dictionary)

	fmt.Println(strings.Join(result, ","))

	dictionaryTree := IndexToBinaryTree(dictionary)

	treeResult := AutoCompleteWithTree(word, dictionaryTree)
	treeResult.Print()

	list := List{nil, nil}

	list.Add("Does")
	list.Add("This")
	list.Add("Work")
	list.Print()

}

func IndexToBinaryTree(dictionary []string) *Node {
	root := Node{
		false,
		"",
		make(map[string]*Node),
	}

	for _, val := range dictionary {
		Index(val, &root)
	}

	return &root
}

func Index(word string, tree *Node) {

	if word == "" {
		return
	}

	firstChar := string(word[0])

	_, ok := tree.Children[firstChar]

	if !ok {

		endOfWord := len(word) == 1

		tree.Children[firstChar] = &Node{
			endOfWord,
			firstChar,
			make(map[string]*Node),
		}
	}

	Index(word[1:], tree.Children[firstChar])
}

type Node struct {
	isEndOfWord bool
	value       string
	Children    map[string]*Node
}

func AutoCompleteWithTree(word string, dictionaryTree *Node) *List {

	wordTree := GetTreeForWord(word, dictionaryTree)

	return GetAllWordsFromTree(wordTree)
}

func GetTreeForWord(word string, dictionaryTree *Node) *Node {

	if word == "" {
		return dictionaryTree
	}

	firstChar := string(word[0])

	val, ok := dictionaryTree.Children[firstChar]

	if ok {
		return GetTreeForWord(word[1:], val)
	}

	return nil
}

func GetAllWordsFromTree(tree *Node) *List {

	list := List{}

	Traverse(tree, &list, "")

	return &list
}

func Traverse(tree *Node, list *List, word string) {

	if tree == nil {
		list.Add(word)
		return
	}

	if tree.isEndOfWord {
		list.Add(word)
	}

	for _, val := range tree.Children {

		newWord := word + tree.value
		Traverse(val, list, newWord)
	}

}

type List struct {
	Head *ListElement
	Tail *ListElement
}

type ListElement struct {
	Value string
	Next  *ListElement
}

func (list *List) Add(value string) {

	element := ListElement{
		value,
		nil,
	}

	if list.Head == nil {
		list.Head = &element
		list.Tail = list.Head
	} else {
		list.Tail.Next = &element
		list.Tail = &element
	}
}

func (list *List) Print() {
	str := ""
	currentElement := list.Head

	for {
		if currentElement == nil {
			break
		}

		str += currentElement.Value + " "

		currentElement = currentElement.Next
	}

	fmt.Println(str)
}

// Complexity is O(mn), where m is the length of the string + n is number of possibilities
func AutoComplete(inputString string, possibilities []string) []string {

	for inputStringIdx, inputChar := range inputString {

		for possibilityIdx, possibility := range possibilities {

			if possibility == "" {
				continue
			}

			if len(possibility) >= (inputStringIdx+1) && inputChar != rune(possibility[inputStringIdx]) {
				possibilities[possibilityIdx] = ""
			}
		}
	}

	possibilities = Shrink(possibilities)

	return possibilities
}

func Shrink(array []string) []string {

	idxValueMap := map[int]string{}

	for _, val := range array {
		if val != "" {
			idxValueMap[len(idxValueMap)] = val
		}
	}

	newArray := make([]string, len(idxValueMap))

	for idx, _ := range newArray {
		newArray[idx] = idxValueMap[idx]
	}

	return newArray
}
