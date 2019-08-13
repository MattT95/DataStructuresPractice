package main

import (
	"fmt"
)

func main() {
	
	type Test struct {
		a int
		b int
	}

	tests := []Test { 
		Test{1, 2},
		Test{2, 3},
		Test{3, 4},
	}

	for _, test  := range tests {
		pair := cons(test.a, test.b)
		first := car(pair)
		second := cdr(pair)

		fmt.Printf("%d - %d", first, second)
		fmt.Println()
	}
}

func cons(a int, b int) func(func(int, int) interface{}) interface{} {
	
	pair := func (f func(int, int) interface{}) interface{} {
		return f(a, b)
	}

	return pair
}

func car(pair func(func(int, int) interface{}) interface{}) int {

	getFirstValue := func(a int, b int) interface{} {
		return interface{}(a)
	}
	
	firstValue, _ := pair(getFirstValue).(int)

	return firstValue
}

func cdr(pair func(func(int, int) interface{}) interface{}) int {

	getFirstValue := func(a int, b int) interface{} {
		return interface{}(b)
	}
	
	firstValue, _ := pair(getFirstValue).(int)

	return firstValue
}