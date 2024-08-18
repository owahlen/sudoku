package main

import (
	"log"
	"os"
	"sudoku/model"
)

func main() {
	// Check if the filename is provided
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}

	filename := os.Args[1]

	board := model.NewBoard()
	err := board.LoadFromFile(filename)
	if err != nil {
		log.Fatalf("Error loading board from file: %v", err)
	}

	// Print the Board
	board.PrintBoard()
}
