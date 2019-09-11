package main

import "fmt"

func main() {

	matrix := [][]int{
		{0, 0, 1},
		{0, 0, 1},
		{1, 0, 0},
	}

	pathCount := FindPathCount(matrix, 0, 0)

	fmt.Println(pathCount)
}

/*
	Could use memoization to reduce computation by caching in regards to position x,y
	(if you know the combinations to that point, we can just take them easily and add on)
*/
func FindPathCount(matrix [][]int, posX int, posY int) int {

	if posX > len(matrix)-1 || posY > len(matrix)-1 {
		return 0
	}

	if matrix[posY][posX] == 1 {
		return 0
	}

	if posX == len(matrix)-1 && posY == len(matrix)-1 {
		return 1
	}

	return FindPathCount(matrix, posX+1, posY) + FindPathCount(matrix, posX, posY+1)
}
