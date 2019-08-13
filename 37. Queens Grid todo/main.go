package main

import (
	"fmt"
)

func main() {

	gridSize := 2

	combinations := queensGrids(gridSize)

	fmt.Println(combinations)
}

// Queen can't be in the same row, column, or diagonal as another queen
// Must be N queens
func queensGrids(gridSize int) int {
	columnsUsed := make([]int, gridSize)

	return Helper(gridSize, columnsUsed, 0)
}

func Helper(gridSize int, columnsUsed []int, column int) int {

	count := 0

	if column == gridSize {
		return 1
	}

	for i := 0; i < gridSize; i++ {

		for j := 0; j < len(columnsUsed); j++ {
			if i == columnsUsed[j] {
				continue
			}

			diagonalIntersectPosInRowOnRight := columnsUsed[j] + (column - j)
			diagonalIntersectPosInRowOnLeft := columnsUsed[j] + (j - column)

			if i == diagonalIntersectPosInRowOnRight || i == diagonalIntersectPosInRowOnLeft {
				continue
			}

			newColumnsUsed := make([]int, len(columnsUsed))
			copy(columnsUsed, newColumnsUsed)
			newColumnsUsed[column] = i

			fmt.Println(i)
			count += Helper(gridSize, newColumnsUsed, column+1)
		}
	}

	return count
}

func PrintGrid(gridSize int, filledPositions []int) {
	for i := 0; i < gridSize; i++ {
		queenPosInRow := filledPositions[i]

		for j := 0; j < gridSize; j++ {
			if queenPosInRow-1 == j {
				fmt.Print("X")
			} else {
				fmt.Print("-")
			}
		}

		fmt.Print("\r\n")
	}
}
