package main

import (
	"fmt"
	"os"
	Quad "sudoku/quad"
)

// ============================================
// HELPER FUNCTION TO CHECK IF A NUMBER CAN BE PLACED (SAFETY CHECK)
// ============================================
func isValid(grid *[9][9]int, row, col, num int) bool {

	// Check the entire ROW and COLUMN
	for i := 0; i < 9; i++ {
		// If number already exists in the row → not valid
		if grid[row][i] == num {
			return false
		}
		// If number already exists in the column → not valid
		if grid[i][col] == num {
			return false
		}
	}

	// Determine which 3×3 box we are in
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3

	// Check all cells in the same 3×3 BOX
	for r := boxRow; r < boxRow+3; r++ {
		for c := boxCol; c < boxCol+3; c++ {
			if grid[r][c] == num {
				return false
			}
		}
	}

	// Passed all checks → safe to place number
	return true
}

// ============================================
// RECURSIVE BACKTRACKING SOLVER
// ============================================
func SolveSudoku(grid *[9][9]int) bool {

	// Search for the next empty cell (represented by 0)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			// Found an empty cell — try numbers 1 to 9
			if grid[row][col] == 0 {

				for num := 1; num <= 9; num++ {

					// Check if placing this number is safe
					if isValid(grid, row, col, num) {

						// Temporarily place the number
						grid[row][col] = num

						// Recursively try to solve the rest
						if SolveSudoku(grid) {
							return true // solved!
						}

						// ❌ Backtrack — undo the placement
						grid[row][col] = 0
					}
				}

				// No valid number found — trigger backtracking
				return false
			}
		}
	}

	// No empty cells left → puzzle solved!
	return true
}

// ============================================
// PRINT THE SUDOKU GRID
// ============================================
func PrintSudoku(grid [9][9]int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Print(grid[row][col])

			// Add a space between numbers
			if col != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// OUR MAIN PROGRAM

func main() {

	Quad.QuadA(5, 3)
	fmt.Println()
	Quad.QuadB(5, 3)
	fmt.Println()
	Quad.QuadC(5, 3)
	fmt.Println()
	Quad.QuadD(5, 3)
	fmt.Println()
	Quad.QuadE(5, 3)
	fmt.Println()

	// Expect exactly 9 rows as input
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	// Create a 9×9 sudoku grid
	var grid [9][9]int

	// Fill the grid from command-line input
	for i := 0; i < 9; i++ {
		row := os.Args[i+1]

		// Each row must have exactly 9 characters
		if len(row) != 9 {
			fmt.Println("Error")
			return
		}

		for j := 0; j < 9; j++ {
			ch := row[j]

			if ch == '.' {
				grid[i][j] = 0 // empty cell
			} else if ch >= '1' && ch <= '9' {
				grid[i][j] = int(ch - '0') // convert ASCII → number
			} else {
				fmt.Println("Error")
				return
			}
		}
	}
	// Solve the puzzle
	if SolveSudoku(&grid) {
		PrintSudoku(grid)
	} else {
		fmt.Println("Error")
	}
}
