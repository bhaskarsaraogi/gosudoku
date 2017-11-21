package main

import "fmt"
import (
	"flag"

	"github.com/bhaskarsaraogi/gosudoku"
)

var p = fmt.Println
const blankSpaces  = 40

func main()  {

	generate := flag.Bool("generate", false, "Genrate a sudoku puzzle")
	solve := flag.Bool("solve", false, "Enter a sudoku puzzle for the system to solve")
	both := flag.Bool("both", false, "Generate a sudoku puzzle and follow with solution")

	var data gosudoku.Grid

	flag.Parse()

	if (*generate && *solve) || *both {
		generateSudoku(&data)

		p("Enter any key to get solution")
		fmt.Scanf("%s")

		solveSudoku(&data)
		return
	}

	if *generate {
		generateSudoku(&data)
	}

	if *solve {
		p("Please enter the sudoku(enter 0 for unassigned location):")

		for i := 0; i < gosudoku.Size; i++ {
			for j := 0; j < gosudoku.Size; j++ {
				fmt.Scanf("%d", &data[i][j])
			}
		}
		p()
		p("SOLUTION:")
		// Solve sudoku
		solveSudoku(&data)
		return
	}

}
func generateSudoku(data *gosudoku.Grid) {
	err := gosudoku.GenerateSudoku(data, blankSpaces)
	if err != nil {
		p(err)
	} else {
		gosudoku.PrintSudoku(data)
	}
	p()
	return
}

func solveSudoku(data *gosudoku.Grid)  {
	if gosudoku.SolveSudoku(data) {
		gosudoku.PrintSudoku(data)
	} else {
		p("No Grid could be formed!")
	}
	return
}
