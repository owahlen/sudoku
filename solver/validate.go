package solver

import "sudoku/model"

// ValidateBoard checks if each digit 1-9 occurs only once in each row and each column
func ValidateBoard(board *model.Board) bool {
	// Check each row
	for row := 0; row < 9; row++ {
		if !validateRow(board, row) {
			return false
		}
	}

	// Check each column
	for column := 0; column < 9; column++ {
		if !validateColumn(board, column) {
			return false
		}
	}

	// Check each 3x3 sub-grid
	for subGridRow := 0; subGridRow < 3; subGridRow++ {
		for subGridColumn := 0; subGridColumn < 3; subGridColumn++ {
			if !validateSubGrid(board, subGridRow, subGridColumn) {
				return false
			}
		}
	}

	return true // No duplicates found
}

func validateRow(board *model.Board, row int) bool {
	seen := make(map[int]bool)
	for column := 0; column < 9; column++ {
		cell := board.Grid[row][column]
		if cell.IsEmpty() {
			continue
		}
		val := *cell.Value
		if seen[val] {
			return false // Duplicate found in row
		}
		seen[val] = true
	}
	return true
}

func validateColumn(board *model.Board, column int) bool {
	seen := make(map[int]bool)
	for row := 0; row < 9; row++ {
		cell := board.Grid[row][column]
		if cell.IsEmpty() {
			continue
		}
		val := *cell.Value
		if seen[val] {
			return false // Duplicate found in row
		}
		seen[val] = true
	}
	return true
}

func validateSubGrid(b *model.Board, subGridRow int, subGridColumn int) bool {
	seen := make(map[int]bool)
	for row := subGridRow * 3; row < subGridRow*3+3; row++ {
		for column := subGridColumn * 3; column < subGridColumn*3+3; column++ {
			cell := b.Grid[row][column]
			if cell.IsEmpty() {
				continue
			}
			val := *cell.Value
			if seen[val] {
				return false // Duplicate found in row
			}
			seen[val] = true
		}
	}
	return true
}
