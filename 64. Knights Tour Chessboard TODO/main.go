package main

// This isn't Knight's Tour currently. See Wiki to understand the problem to solve

import (
	"fmt"
)

func main() {

	boardSize := 4

	startPosX := 2
	startPosY := 2

	usedMoves := make([][]int, 4, 4)
	moves := CountMoves(boardSize, startPosX, startPosY, usedMoves)

	fmt.Println(moves)
}

func CountMoves(boardBound int, posX int, posY int, usedMoves [][]int) int {

	fmt.Printf("%d - %d", posX, posY)
	fmt.Println()

	count := 0

	usedMoves[posX][posY] = 1

	count += CountOnMove(boardBound, posX, posY, -1, 2)
	count += CountOnMove(boardBound, posX, posY, 1, 2)

	count += CountOnMove(boardBound, posX, posY, -1, -2)
	count += CountOnMove(boardBound, posX, posY, 1, -2)

	count += CountOnMove(boardBound, posX, posY, 2, 1)
	count += CountOnMove(boardBound, posX, posY, 2, -1)

	count += CountOnMove(boardBound, posX, posY, -2, 1)
	count += CountOnMove(boardBound, posX, posY, -2, -1)

	return count
}

func CountOnMove(boardBound int, posX int, posY int, transitionX int, transitionY int) int {

	newPosX := posX + transitionX
	newPosY := posY + transitionY

	if newPosX > posX && newPosX >= boardBound {
		return 0
	}

	if newPosY > posY && newPosY >= boardBound {
		return 0
	}

	if newPosX <= posX && newPosX < 0 {
		return 0
	}

	if newPosY <= posY && newPosY < 0 {
		return 0
	}

	return 1 + CountMoves(boardBound, newPosX, newPosY)
}
