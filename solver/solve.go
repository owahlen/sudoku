package solver

import (
	"errors"
	"sudoku/model"
	"sync"
)

const unsolvableBoard string = "unsolvable board"

// SolveBoard solves the sudoku board or returns an error
func SolveBoard(board *model.Board) (*model.Board, error) {

	// check if the board is valid
	if !ValidateBoard(board) {
		return nil, errors.New(unsolvableBoard)
	}

	// find the first cell that is empty
	row, column, found := findFirstEmptyCell(board)
	if !found {
		// the board is valid with no empty cells means the solution is found
		return board, nil
	}

	// Channel to receive solutions from goroutines
	solutionChan := make(chan *model.Board)
	errorChan := make(chan error)
	var wg sync.WaitGroup

	// try all digits in the empty cell and recurse to find the solution
	for i := 1; i <= 9; i++ {
		wg.Add(1)
		go func(digit int) {
			defer wg.Done()
			solution, err := tryDigit(board, row, column, digit)
			if err != nil {
				errorChan <- err
			} else {
				solutionChan <- solution
			}
		}(i)
	}

	// Goroutine to close the channels when all goroutines are done
	go func() {
		wg.Wait()
		close(solutionChan)
		close(errorChan)
	}()

	// Wait for a solution or return an error if no solution is found
	for {
		select {
		case solution, ok := <-solutionChan:
			if ok {
				return solution, nil
			} else {
				// solution channel was closed without finding a solution
				return nil, errors.New(unsolvableBoard)
			}
		case err, ok := <-errorChan:
			// ignore unsolvable board errors from the recursive calls
			if ok && err.Error() != unsolvableBoard {
				return nil, err
			}
		}
	}
}

func tryDigit(board *model.Board, row, column, value int) (*model.Board, error) {
	clone := board.Clone()
	err := clone.SetCell(row, column, value)
	if err != nil {
		return nil, err
	}
	solution, err := SolveBoard(clone)
	if err == nil {
		// the attempt was successful:
		return solution, err
	}
	clone = nil
	return nil, errors.New("unsolvable board")
}

func findFirstEmptyCell(board *model.Board) (row, column int, found bool) {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if board.Grid[row][column] == 0 {
				return row, column, true
			}
		}
	}
	return -1, -1, false
}
