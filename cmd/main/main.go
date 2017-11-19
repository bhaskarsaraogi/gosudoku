package main

import "fmt"
import "github.com/bhaskarsaraogi/gosudoku"

func main()  {

	data := gosudoku.Grid{}
	if gosudoku.SolveSudoku(&data) {
		gosudoku.PrintSudoku(&data)
	} else {
		fmt.Println("No Grid could be formed!")
	}

}
