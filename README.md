gosudoku is a command line util to generate sudoku puzzle, solve sudoku or both

## Installation

```
$ go get github.com/bhaskarsaraogi/gosudoku/cmd/gosudoku
```

## Play

####Generate a sudoku puzzle:
```
$ ./gosudoku -generate
```
######Output

    9 6 0 1 0 0 5 7 8 
    0 5 0 3 7 9 0 0 6 
    0 7 2 0 0 8 9 1 3 
    0 0 0 0 0 3 8 5 0 
    2 9 0 8 0 7 0 0 0 
    1 0 8 9 0 0 0 0 0 
    3 8 4 0 0 5 1 0 0 
    0 0 6 4 8 0 3 0 0 
    5 2 0 6 0 1 0 8 0 


####Solve a sudoku puzzle(will be asked to enter sudoku to solve):
```
$ ./gosudoku -solve
```
###### Output
```
Please enter the sudoku(enter 0 for unassigned location):
4 9 2 5 1 3 7 0 6 
0 0 0 2 0 8 3 0 1 
8 0 0 7 0 0 0 2 0 
0 0 5 0 6 0 0 3 0 
0 7 8 3 0 0 0 0 0 
3 6 9 0 8 7 5 1 2 
0 0 4 9 0 0 0 0 0 
0 3 0 8 0 1 6 0 0 
9 8 0 6 7 4 0 5 0 
 
SOLUTION:
 
4 9 2 5 1 3 7 8 6 
7 5 6 2 4 8 3 9 1 
8 1 3 7 9 6 4 2 5 
2 4 5 1 6 9 8 3 7 
1 7 8 3 5 2 9 6 4 
3 6 9 4 8 7 5 1 2 
6 2 4 9 3 5 1 7 8 
5 3 7 8 2 1 6 4 9 
9 8 1 6 7 4 2 5 3 

```

####Both(will be asked to press a key to retrieve solution):
```
$ ./gosudoku -both
```

###### Output

```
7 0 0 0 0 5 9 0 8 
2 1 9 3 4 0 0 7 0 
8 5 0 0 0 0 2 3 1 
0 0 0 5 1 3 7 0 9 
5 0 0 8 7 4 3 0 6 
0 0 8 6 0 0 0 5 0 
0 0 2 0 0 0 0 0 0 
6 8 5 2 3 1 0 9 0 
0 3 0 4 0 0 8 1 0 
 
Enter any key to get solution
 
7 6 3 1 2 5 9 4 8 
2 1 9 3 4 8 6 7 5 
8 5 4 7 6 9 2 3 1 
4 2 6 5 1 3 7 8 9 
5 9 1 8 7 4 3 2 6 
3 7 8 6 9 2 1 5 4 
1 4 2 9 8 7 5 6 3 
6 8 5 2 3 1 4 9 7 
9 3 7 4 5 6 8 1 2 

```