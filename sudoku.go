package main

import (
    "os"

    "github.com/01-edu/z01"
)

var puzzle [9][9]rune // A 2D array with dimensions 9x9

/* BoardValidity  function checks if a given number can be placed at a specific position in the puzzle */

func BoardValidity(guess rune, row, col int) bool {
    for j := 0; j < 9; j++ {
        if puzzle[row][j] == guess {
            return false
        }
    }
    for i := 0; i < 9; i++ {
        if puzzle[i][col] == guess {
            return false
        }
    }
    row_start := (row / 3) * 3
    col_start := (col / 3) * 3
    for i := row_start; i < row_start+3; i++ {
        for j := col_start; j < col_start+3; j++ {
            if puzzle[i][j] == guess {
                return false
            }
        }
    }
    return true
}

/* NextEmpty function finds the indices of the next empty cell in the puzzle*/

func NextEmpty() (int, int) {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if puzzle[i][j] == '.' {
                return i, j
            }
        }
    }
    return -1, -1
}

/*SudokuSolver function uses a recursive backtracking algorithm to solve the Sudoku puzzle. It fills in the empty cells one by one, checking if each guess is valid and backtracking if it leads to an invalid solution.*/

func SudokuSolver() bool {
    row, col := NextEmpty()
    if row == -1 {
        return true
    }
    for i := 1; i < 10; i++ {
        if BoardValidity(rune(i+48), row, col) == true {
            puzzle[row][col] = rune(i + 48)
            if SudokuSolver() == true {
                return true
            }
        }
        puzzle[row][col] = '.'
    }
    return false
}

/*ReportError function is called when an invalid input is detected*/

func ReportError() {
    error := "Error"
    for i := 0; i < len(error); i++ {
        z01.PrintRune(rune(error[i]))
    }
    z01.PrintRune('\n')
}

/*In the main function, we read the command-line arguments provided by the user. 
We validate the input by checking if there are exactly 9 arguments and each argument has a length of 9.
We loop through each argument and each character in the argument. 
If the character is a digit from 1 to 9, 
we check if it is valid for the current position using the BoardValidity function. 
If it is valid, we store it in the puzzle array. If the character is a period (.), 
it represents an empty cell, so we store a period in the puzzle array. 
If the character is neither a digit nor a period, we print an error message and exit the program.
If the input is valid, we call the SudokuSolver function to solve the puzzle.
If a solution is found, we print the solved puzzle by looping through the puzzle array and printing each digit or period.
We separate each digit by a space, and each row is printed on a new line.
If a solution is not found, we print an error message.*/

func main() {
    args := os.Args[1:]
    if len(args) != 9 {
        ReportError()
        return
    }
    for i := 0; i < 9; i++ {
        if len(args[i]) != 9 {
            ReportError()
            return
        }
        for j := 0; j < 9; j++ {
            if args[i][j] > '0' && args[i][j] <= '9' {
                if BoardValidity(rune(args[i][j]), i, j) == true {
                    puzzle[i][j] = rune(args[i][j])
                } else {
                    ReportError()
                    return
                }
            } else {
                if args[i][j] != '.' {
                    ReportError()
                    return
                }
                puzzle[i][j] = '.'
            }
        }
    }
    if SudokuSolver() == true {
        for i := 0; i < 9; i++ {
            for j := 0; j < 9; j++ {
                z01.PrintRune(puzzle[i][j])
                if j != 8 {
                    z01.PrintRune(' ')
                }
            }
            z01.PrintRune('\n')
        }
    } else {
        ReportError()
        return
    }
}
