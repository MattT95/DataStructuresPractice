package ReverseLinkedList

func CreateTestLinkedList() LinkedListElement {
	
}

func ReverseLinkedList(start LinkedListElement) LinkedListElement {
	return LinkedListElement()
}



type LinkedListElement struct {
	Value int
	Next *LinkedListElement

}

type LinkedList struct {
	head *LinkedListElement
	tail *LinkedListElement
}

func (list LinkedList) Add(value int) {
	
	if list.head == nil {
		element := LinkedListElement{ value, nil }
	}

	if list.tail == nil {
		element := LinkedListElement{ value, nil }
		list.head.Next = &element
		list.tail = &element
	}

	element := LinkedListElement{ value, nil }
	list.tail.Next = &element
}

func (list LinkedList) Print() string {
	currentNode := *list.head


	
}

