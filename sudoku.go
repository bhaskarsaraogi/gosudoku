package gosudoku

import (
	"fmt"
	//"math/rand"
)

type Grid [9][9]int

func SolveSudoku(data *Grid) bool {

	var row, col int

	if !findUnassignedLocation(data, &row, &col) {
		return true
	}

	for i:=1; i<=9; i++ {
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

func findUnassignedLocation(data *Grid, row *int, col *int) bool{
	for *row = 0; *row < 9; *row += 1 {
		for *col = 0; *col < 9; *col += 1 {
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
	for i:=0; i<9; i++ {
		if data[row][i]==value {
			return false
		}
	}
	return true
}

func isValidInCol(data *Grid, col int, value int) bool {
	for i:=0; i<9; i++ {
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
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(data[i][j], " ")
		}
		fmt.Println()
	}
}
