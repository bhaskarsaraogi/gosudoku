package main

import "fmt"
import (
	"github.com/bhaskarsaraogi/gosudoku"
	"flag"
)

func main()  {

	generate := flag.Bool("generate", false, "Genrate a sudoku puzzle")
	solve := flag.Bool("solve", false, "Enter a sudoku puzzle for the system to solve")
	both := flag.Bool("both", false, "Generate a sudoku puzzle and follow with solution")

	var data gosudoku.Grid
	defer fmt.Println()

	flag.Parse()

	if (*generate && *solve) || *both {
		err := gosudoku.GenerateSudoku(&data, 40)
		if err != nil {
			fmt.Println(err)
		} else {
			gosudoku.PrintSudoku(&data)
		}

		fmt.Println()

		fmt.Println("Enter any key to get solution")
		fmt.Scanf("%s")

		if gosudoku.SolveSudoku(&data) {
			gosudoku.PrintSudoku(&data)
		} else {
			fmt.Println("No Grid could be formed!")
		}
		return
	}

	if *generate {
		err := gosudoku.GenerateSudoku(&data, 40)
		if err != nil {
			fmt.Println(err)
		} else {
			gosudoku.PrintSudoku(&data)
		}
		return
	}

	if *solve {
		fmt.Println("Please enter the sudoku(enter 0 for unassigned location):")

		for i:=0; i<9; i++ {
			for j:=0; j<9; j++ {
				fmt.Scanf("%d", &data[i][j])
			}
		}

		fmt.Println()

		// Solve sudoku
		if gosudoku.SolveSudoku(&data) {
			gosudoku.PrintSudoku(&data)
		} else {
			fmt.Println("No Grid could be formed!")
		}
		return
	}

}
