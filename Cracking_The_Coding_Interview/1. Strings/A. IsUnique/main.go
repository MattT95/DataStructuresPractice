package main

import (
	"fmt"
)

/*
	Implement an algorithm to determine if a string has all unique characters.
	What if you can't use additional data structures?
*/

func main() {

	/*
		One solution is to use a data structure, specifically a hash table.

		Whilst traversing through the string, we can check the hash table to see whether we
		have encountered a specific key yet. If we have, then we know that there is a duplicate character,
		and therefore we should return False to indicate that the characters are not unique.
	*/

	uniqueTestCase := "abcdefghijklmopqrstuvwxyz"
	nonUniqueTestCase := "abcdefghijklmnopa"

	fmt.Println(IsUniqueWithDataStruture(uniqueTestCase))
	fmt.Println(IsUniqueWithDataStruture(nonUniqueTestCase))
}

func IsUniqueWithDataStruture(str string) bool {

	table := NewHashTable(8)

	for i := 0; i < len(str); i++ {
		value := str[i]
		fmt.Println(rune(value))

		if table.Contains(rune(value)) {
			return false
		}

		table.Add(rune(value))
	}

	return true
}

type HashTable struct {
	table []*LinkedList
}

type Entry struct {
	value rune
}

func NewHashTable(size int) HashTable {
	return HashTable{
		table: make([]*LinkedList, size),
	}
}

func (hashtable *HashTable) Add(key rune) {
	hash := int(key) % len(hashtable.table)

	if hashtable.table[hash] == nil {
		hashtable.table[hash] = &LinkedList{}
	}

	list := hashtable.table[hash]

	list.Add(key)
}

func (hashtable *HashTable) Contains(key rune) bool {

	hash := int(key) % len(hashtable.table)

	list := hashtable.table[hash]

	if list != nil && list.Contains(key) {
		return true
	}
	return false
}

type LinkedList struct {
	first *LinkedListNode
	last  *LinkedListNode
}

type LinkedListNode struct {
	value rune
	next  *LinkedListNode
}

func (linkedList *LinkedList) Add(val rune) {
	node := LinkedListNode{
		value: val,
		next:  nil,
	}

	if linkedList.first == nil {
		linkedList.first = &node
		linkedList.last = &node
	} else {
		linkedList.last.next = &node
		linkedList.last = &node
	}
}

func (linkedList *LinkedList) Contains(val rune) bool {
	nextNodeToCheck := linkedList.first

	for nextNodeToCheck != nil {
		if (*nextNodeToCheck).value == val {
			return true
		}

		nextNodeToCheck = (*nextNodeToCheck).next
	}

	return false
}
