package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Board represents a 9x9 grid of Cells
type Board struct {
	Grid [9][9]int `json:"grid"` // Use JSON tags to define the JSON structure
}

// NewBoard creates a new empty Board
func NewBoard() *Board {
	return &Board{}
}

// SetCell sets a cell in the Board to a specific value (1-9)
func (board *Board) SetCell(row, column, value int) error {
	if row < 0 || row > 8 || column < 0 || column > 8 {
		return errors.New("row and column must be between 0 and 8")
	}
	if value < 0 || value > 9 {
		return errors.New("value must be between 0 and 9")
	}
	board.Grid[row][column] = value
	return nil
}

// GetCell returns the value of a cell or an error if the cell is empty
func (board *Board) GetCell(row, column int) (int, error) {
	if row < 0 || row > 8 || column < 0 || column > 8 {
		return 0, errors.New("row and column must be between 0 and 8")
	}
	return board.Grid[row][column], nil
}

// PrintBoard prints the Board to the console
func (board *Board) PrintBoard() {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if board.Grid[row][column] == 0 {
				fmt.Print(". ") // Print a dot for empty cells
			} else {
				value := board.Grid[row][column]
				fmt.Printf("%d ", value)
			}
		}
		fmt.Println()
	}
}

// LoadFromFile loads the Board from a JSON file
func (board *Board) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return board.LoadFromJSON(data)
}

// LoadFromJSON loads the Board from a JSON string
func (board *Board) LoadFromJSON(data []byte) error {
	return json.Unmarshal(data, board)
}

// Clone creates a deep copy of the Board
func (board *Board) Clone() *Board {
	clone := NewBoard()
	// Grid consist of primitive types and therefore is copied by value
	clone.Grid = board.Grid
	return clone
}
