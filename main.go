package main

import (
	"sudoku/model"
)

func main() {
	board := model.NewBoard()

	// Example: Set some cells to specific values
	board.SetCell(0, 0, 5)
	board.SetCell(1, 1, 3)
	board.SetCell(2, 2, 6)
	board.SetCell(3, 3, 7)
	board.SetCell(4, 4, 1)
	board.SetCell(5, 5, 9)
	board.SetCell(6, 6, 5)
	board.SetCell(7, 7, 2)
	board.SetCell(8, 8, 4)

	// Print the Board
	board.PrintBoard()
}
