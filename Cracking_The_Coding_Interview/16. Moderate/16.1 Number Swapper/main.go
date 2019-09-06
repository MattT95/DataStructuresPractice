package main

import "fmt"

/*
	Number Swapper
		Write a function to swap a number in place (that is, without temporary variables).
*/

func main() {

	num1 := 5
	num2 := 8

	Swap(&num1, &num2)

	fmt.Println(num1)
	fmt.Println(num2)
}

func Swap(num1 *int, num2 *int) {
	*num1 = *num1 ^ *num2
	*num2 = *num2 ^ *num1 // since num1 is an XOR of both number, num2 is "cancelled" out
	*num1 = *num1 ^ *num2 // and since num2 is num1, XOR cancels out num1
}
