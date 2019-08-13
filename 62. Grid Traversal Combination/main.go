package main

import "fmt"

func main() {

	matrixSizeX := 50
	matrixSizeY := 50

	combinationsA := CountNumberOfTraversals(matrixSizeX, matrixSizeY, 0, 0)
	fmt.Println(combinationsA)

	memo := make([][]int, matrixSizeX)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, matrixSizeY)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}

	// combinationsB := CountNumberOfTraversalsMemo(matrixSizeX, matrixSizeY, memo)
	// fmt.Println(combinationsB)
}

func CountNumberOfTraversals(matrixSizeX int, matrixSizeY int, posX int, posY int) int {

	if posX == matrixSizeX-1 && posY == matrixSizeY-1 {
		return 1
	}

	count := 0

	if posX != matrixSizeX-1 {
		count += CountNumberOfTraversals(matrixSizeX, matrixSizeY, posX+1, posY)
	}

	if posY != matrixSizeY-1 {
		count += CountNumberOfTraversals(matrixSizeX, matrixSizeY, posX, posY+1)
	}

	return count
}

func CountNumberOfTraversalsMemo(matrixSizeX int, matrixSizeY int, memo [][]int) int {

	if matrixSizeX-1 == 0 && matrixSizeY-1 == 0 {
		return 1
	}

	if memo[matrixSizeX-1][matrixSizeY-1] != -1 {
		return memo[matrixSizeX-1][matrixSizeY-1]
	}

	if memo[matrixSizeY-1][matrixSizeX-1] != -1 {
		return memo[matrixSizeY-1][matrixSizeX-1]
	}

	count := 0

	if matrixSizeX-1 != 0 {
		count += CountNumberOfTraversalsMemo(matrixSizeX-1, matrixSizeY, memo)
	}

	if matrixSizeY-1 != 0 {
		count += CountNumberOfTraversalsMemo(matrixSizeX, matrixSizeY-1, memo)
	}

	memo[matrixSizeX-1][matrixSizeY-1] = count
	memo[matrixSizeY-1][matrixSizeX-1] = count

	return count
}
