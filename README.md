gosudoku is a command line util to generate sudoku puzzle, solve sudoku or both

## Installation

```
$ go get github.com/bhaskarsaraogi/gosudoku/cmd/gosudoku
```

## Play

Generate a sudoku puzzle:
```
$ ./gosudoku -generate
```

Solve a sudoku puzzle(will be asked to enter sudoku to solve):
```
$ ./gosudoku -solve
```

Both(will be asked to press a key to retrieve solution):
```
$ ./gosudoku -both
```