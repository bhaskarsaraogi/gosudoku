package gosudoku

import (
	"fmt"
	"math/rand"
	"errors"
	"time"
	"math"
)

// Size of sudoku
const Size = 9

// Define the grid of 9X9 as a type
type Grid [Size][Size]int


// Solve the sudoku grid, return true if solved else false
// fixme pass by value not refrence, if couldnt be solved we should be in a position to trace back, or is it already done because of backtracking ?
func SolveSudoku(data *Grid) bool {

	var row, col int

	// find an unassigned location in the grid, if none then sudoku is solved return true
	if !findUnassignedLocation(data, &row, &col) {
		return true
	}

	// try out values between 1-9 one at a time, if works recurse to solve further if none works then sudoku cannot be solved
	for i:=1; i<= Size; i++ {
		if isValid(data, row, col, i) {

			// since valid currently, mark this guy with i
			data[row][col] = i

			// solve ahead
			if SolveSudoku(data) {
				return true
			} else {
				// i didnt work, we will try next number so backtrack and mark is at not solved
				data[row][col] = 0
			}
		}
	}

	return false
}


// Generate a random Sudoku puzzle, parama data is the sudoku grid, emptyCells is the number of empty cells we need in the puzzle
func GenerateSudoku(data *Grid, emptyCells int) error {

	g := 0
	// to be used as a marker that all numbers 1-9 have been assigned in a 3X3 grid
	set := make(map[int]bool)

	// seed random number generator with current system time, to add jitter
	rand.Seed(time.Now().Unix())

	// fill diagonal of the grid, i.e the 3 3X3 grids
	for ; g < 3 ; g++ {
		for i:=0 + g*3; i<3 + g*3; i++ {
			for j:=0 + g*3; j < 3 + g*3; j++ {

				// generate a random number between 1-9
				n := rand.Intn(Size)+1
				// see if it has been assigned for this 3X3 grid, if yes, choose another(for loop)
				for set[n] {
					n = rand.Intn(Size)+1
				}

				// finally we have a number which has not been assigned in this 3X3 grid
				data[i][j] = n //assign
				set[n] = true // claim assignment
			}
		}
		// reset the markers to start processing next grid
		set = make(map[int]bool)
		// traverse over the next diagonal(left to right, towards bottom) 3X3 grid
	}

	// Now all the diagonal(left to right, towards bottom) 3X3 grids, has been filled
	// Just solve this grid as a solution and we have a complete solved sudoku
	if !SolveSudoku(data) {
		return errors.New("Unable to solve sudoku")
	}

	// emptyCells from generated sudoku, to make it a puzzle
	for emptyCells > 0 {

		// choose random index, take modulo with Size so that both i,j are betweeb 0,0-8,8
		// mark them empty
		if i,j := rand.Intn(math.MaxInt32), rand.Intn(math.MaxInt32); data[i%Size][j%Size] != 0 {
			data[i%Size][j%Size] = 0
			emptyCells--
		}
	}

	return nil
}

// Find an unassigned location and pass by reference in row & col variables the value of cell which is empty, else return false(no cell is empty anymore)
func findUnassignedLocation(data *Grid, row *int, col *int) bool{
	for *row = 0; *row < Size; *row += 1 {
		for *col = 0; *col < Size; *col += 1 {
			if data[*row][*col] == 0 {
				return true
			}
		}
	}
	return false
}

// check if a particular value fits in particular position, in the row, col & grid
func isValid(data *Grid, row int, col int, value int) bool {
	if isValidInRow(data, row, value) && isValidInCol(data, col, value) && isValidInGrid(data, row, col, value) {
		return true
	}
	return false
}

// check if a value fits in row or is already assigned
func isValidInRow(data *Grid, row int, value int) bool {
	for i:=0; i< Size; i++ {
		if data[row][i]==value {
			return false
		}
	}
	return true
}

// check if a value fits in col or is already assigned
func isValidInCol(data *Grid, col int, value int) bool {
	for i:=0; i< Size; i++ {
		if data[i][col]==value {
			return false
		}
	}
	return true
}

// check if a value fits in grid or is already assigned
func isValidInGrid(data *Grid, row int, col int, value int) bool {
	for i:=row - row%3; i < row - row%3 + 3; i++ {
		for j:=col - col%3; j < col - col%3 + 3; j++ {
			if data[i][j] == value {
				return false
			}
		}
	}
	return true
}

// print sudoku
func PrintSudoku(data *Grid) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			fmt.Print(data[i][j], " ")
		}
		fmt.Println()
	}
}
