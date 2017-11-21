package gosudoku

import (
	"fmt"
	"math/rand"
	"errors"
	"time"
	"math"
)

const Size = 9

type Grid [Size][Size]int

func SolveSudoku(data *Grid) bool {

	var row, col int

	if !findUnassignedLocation(data, &row, &col) {
		return true
	}

	for i:=1; i<= Size; i++ {
		if isValid(data, row, col, i) {

			data[row][col] = i

			if SolveSudoku(data) {
				return true
			} else {
				data[row][col] = 0
			}
		}
	}

	return false
}

func GenerateSudoku(data *Grid, removeNDigits int) error {
	// implement generation logic
	g := 0
	set := make(map[int]bool)
	rand.Seed(time.Now().Unix())

	for g < 3 {
		for i:=0 + g*3; i<3 + g*3; i++ {
			for j:=0 + g*3; j < 3 + g*3; j++ {

				n := rand.Intn(Size)+1
				for set[n] {
					n = rand.Intn(Size)+1
				}
				data[i][j] = n
				set[n] = true
			}
		}
		g++
		set = make(map[int]bool)
	}

	if !SolveSudoku(data) {
		return errors.New("Unable to solve sudoku")
	}

	// removeNDigits from generated sudoku
	for removeNDigits > 0 {

		if i,j := rand.Intn(math.MaxInt32), rand.Intn(math.MaxInt32); data[i%Size][j%Size] != 0 {
			data[i%Size][j%Size] = 0
			removeNDigits--
		}
	}

	return nil
}

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

func isValid(data *Grid, row int, col int, value int) bool {
	if isValidInRow(data, row, value) && isValidInCol(data, col, value) && isValidInGrid(data, row, col, value) {
		return true
	}
	return false
}

func isValidInRow(data *Grid, row int, value int) bool {
	for i:=0; i< Size; i++ {
		if data[row][i]==value {
			return false
		}
	}
	return true
}

func isValidInCol(data *Grid, col int, value int) bool {
	for i:=0; i< Size; i++ {
		if data[i][col]==value {
			return false
		}
	}
	return true
}

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

func PrintSudoku(data *Grid) {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			fmt.Print(data[i][j], " ")
		}
		fmt.Println()
	}
}
