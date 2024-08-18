package solver

import (
	"errors"
	"sudoku/model"
)

// SolveBoard solves the sudoku board or returns an error
func SolveBoard(board *model.Board) (solvedBoard *model.Board, err error) {

	// check if the board is valid
	if !ValidateBoard(board) {
		return nil, errors.New("unsolvable board")
	}

	// find the first cell that is empty
	row, column, found := findFirstEmptyCell(board)
	if !found {
		// the board is valid with no empty cells means the solution is found
		return board, nil
	}

	// try all digits in the empty cell and recurse to find the solution
	for i := 1; i < 10; i++ {
		clone := board.Clone()
		err := clone.SetCell(row, column, i)
		if err != nil {
			return nil, err
		}
		solution, err := SolveBoard(clone)
		if err == nil {
			// the attempt was successful:
			return solution, err
		}
		clone = nil
	}
	return nil, errors.New("unsolvable board")
}

func findFirstEmptyCell(board *model.Board) (row, column int, found bool) {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if board.Grid[row][column].IsEmpty() {
				return row, column, true
			}
		}
	}
	return -1, -1, false
}
