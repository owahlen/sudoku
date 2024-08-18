package main

import (
	"log"
	"sudoku/model"
)

func main() {
	boardJson := getBoardJsonString()
	board := model.NewBoard()

	err := board.LoadFromJSON([]byte(boardJson))
	if err != nil {
		log.Fatalf("Error loading board from JSON: %v", err)
	}

	// Print the Board
	board.PrintBoard()
}

func getBoardJsonString() string {
	// Example JSON data representing a 9x9 board
	jsonData := `
	{
		"grid": [
			[5, 3, null, null, 7, null, null, null, null],
			[6, null, null, 1, 9, 5, null, null, null],
			[null, 9, 8, null, null, null, null, 6, null],
			[8, null, null, null, 6, null, null, null, 3],
			[4, null, null, 8, null, 3, null, null, 1],
			[7, null, null, null, 2, null, null, null, 6],
			[null, 6, null, null, null, null, 2, 8, null],
			[null, null, null, 4, 1, 9, null, null, 5],
			[null, null, null, null, 8, null, null, 7, 9]
		]
	}
	`
	return jsonData
}
