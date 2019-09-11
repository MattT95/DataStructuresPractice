package main

import(
	"fmt"
)

func main() {
	// 1111 0000 1111 0000 1111 0000 1111 0000
	val := 0B_11110000

	val = val ^ 0B_11111111

	fmt.Printf("%b", val)
}