package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Board represents a 9x9 grid of Cells
type Board struct {
	grid [9][9]Cell `json:"grid"` // Use JSON tags to define the JSON structure
}

// NewBoard creates a new empty Board
func NewBoard() *Board {
	return &Board{}
}

// SetCell sets a cell in the Board to a specific value (1-9)
func (b *Board) SetCell(row, col, value int) error {
	if row < 0 || row >= 9 || col < 0 || col >= 9 {
		return errors.New("row and column must be between 0 and 8")
	}
	return b.grid[row][col].SetValue(value)
}

// ClearCell clears a cell in the Board, making it empty
func (b *Board) ClearCell(row, col int) error {
	if row < 0 || row >= 9 || col < 0 || col >= 9 {
		return errors.New("row and column must be between 0 and 8")
	}
	b.grid[row][col].Clear()
	return nil
}

// GetCell returns the value of a cell or an error if the cell is empty
func (b *Board) GetCell(row, col int) (int, error) {
	if row < 0 || row >= 9 || col < 0 || col >= 9 {
		return 0, errors.New("row and column must be between 0 and 8")
	}
	return b.grid[row][col].GetValue()
}

// PrintBoard prints the Board to the console
func (b *Board) PrintBoard() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.grid[i][j].IsEmpty() {
				fmt.Print(". ") // Print a dot for empty cells
			} else {
				val, _ := b.grid[i][j].GetValue()
				fmt.Printf("%d ", val)
			}
		}
		fmt.Println()
	}
}

// LoadFromJSON loads the Board from a JSON string
func (b *Board) LoadFromJSON(data []byte) error {
	return json.Unmarshal(data, b)
}
