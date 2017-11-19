package main

import "fmt"
import "github.com/bhaskarsaraogi/gosudoku"

func main()  {

	var data gosudoku.Grid

	err := gosudoku.GenerateSudoku(&data, 40)
	if err != nil {
		fmt.Println(err)
	} else {
		gosudoku.PrintSudoku(&data)
	}

	fmt.Println()
	defer fmt.Println()

	if gosudoku.SolveSudoku(&data) {
		gosudoku.PrintSudoku(&data)
	} else {
		fmt.Println("No Grid could be formed!")
	}

}
