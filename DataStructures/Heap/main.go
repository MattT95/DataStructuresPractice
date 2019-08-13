package main

import "fmt"

func main() {

	array := []int{24, 30, 6, 7, 1, 55, 78}
	heap := CreateHeap(&array)

	heap.Print()
}

func CreateHeap(array *[]int) Heap {
	h := Heap{*array}

	h.Heapify()

	return h
}

type Heap struct {
	values []int
}

func (h *Heap) Add(value int) {
	fmt.Println("do nothing")
}

func (h *Heap) Heapify() {
	Heapify(&h.values)
}

func Heapify(array *[]int) {
	for i := (len(*array) / 2) - 1; i >= 0; i-- {
		SiftDown(array, i)
	}
}

func SiftDown(array *[]int, rootIdx int) {
	firstChildIdx := (2 * rootIdx) + 1
	secondChildIdx := (2 * rootIdx) + 2

	largestIdx := rootIdx

	fmt.Printf("Root: %d, Left_Child: %d, Right_Child: %d \n", rootIdx, firstChildIdx, secondChildIdx)

	if firstChildIdx <= len(*array)-1 && (*array)[firstChildIdx] > (*array)[largestIdx] {
		largestIdx = firstChildIdx
	}

	if secondChildIdx <= len(*array)-1 && (*array)[secondChildIdx] > (*array)[largestIdx] {
		largestIdx = secondChildIdx
	}

	if rootIdx != largestIdx {
		fmt.Printf("Swapping %d with %d \n", (*array)[rootIdx], (*array)[largestIdx])
		temp := (*array)[rootIdx]
		(*array)[rootIdx] = (*array)[largestIdx]
		(*array)[largestIdx] = temp
		SiftDown(array, largestIdx)
	}
}

func (h *Heap) Print() {
	fmt.Print("[")
	for i := 0; i < len(h.values); i++ {
		fmt.Print(" ")
		fmt.Print(h.values[i])
	}
	fmt.Print(" ]")
}
