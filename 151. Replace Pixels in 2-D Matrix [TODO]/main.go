package main

/*
* Given the inputs:
*	- a 2-D Matrix representing an image, a l
*	- a location of a pixel in the screen
*	- a color C
*
* Replace the colour of a given pixel and all the adjacent same coloured pixels with C
*
* For example, given the following matrix and location (2, 2) and 'G' for green:
*
*	B	B	W				B	B	G
*	W	W	W	becomes		G	G	G
*	W	W	W				G	G	G
*	B	B	B				B	B	B
*
 */

func main() {
	input := [][]rune{
		{'B', 'B', 'W'},
		{'W', 'W', 'W'},
		{'W', 'W', 'W'},
		{'B', 'B', 'B'},
	}

	position := Point{2, 2}
	positionY := 2

	colour := 'G'

	ReplaceColour(&input, position, colour)

	Print(input)
}

func Print([][]int matrix){
	for i := 0; i < len(matrix); i++ {
		row := matrix[i]

		for j := 0; j < len(row); j++ {
			string.jon
		}

	}
}
type Point struct {
	X int
	Y int
}
