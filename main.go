package main

import (
	"fmt"
	"log"
	"os"
	"sudoku/model"
	"sudoku/solver"
)

func main() {
	// Check if the filename is provided
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}
	filename := os.Args[1]

	// load the initial board
	board := model.NewBoard()
	err := board.LoadFromFile(filename)
	if err != nil {
		log.Fatalf("Error loading board from file: %v", err)
	}
	fmt.Println("Solving board:")
	board.PrintBoard()

	// solve the board
	solution, err := solver.SolveBoard(board)

	// Print the Board
	if err != nil {
		log.Fatalf("Error solving board: %v", err)
	} else {
		fmt.Println("Solution:")
		solution.PrintBoard()
	}

}
