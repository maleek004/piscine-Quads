package Quad

import "fmt" // Borrow the printing tool

// ============================================
// TOOL #1: THE SAFETY CHECKER
// Job: Check if we can put a number in a spot
// ============================================
func isValid(grid *[9][9]int, row, col, num int) bool {
	// grid = the 9x9 puzzle board
	// row = which row we're checking (0-8)
	// col = which column we're checking (0-8)
	// num = the number we want to place (1-9)
	// Returns: true (safe!) or false (not safe!)

	// STEP 1: Check the ROW (going sideways ↔️)
	for i := 0; i < 9; i++ {
		// Look at every spot in this row
		// Is the number already there?
		if grid[row][i] == num || grid[i][col] == num {
			// Found it in the row OR column!
			return false // Not allowed! 🚫
		}
	}

	// STEP 2: Check the 3×3 BOX (📦)
	// First, find which box we're in
	startRow := (row / 3) * 3 // Find box's starting row
	startCol := (col / 3) * 3 // Find box's starting column

	// Example: If row=5, col=7
	// startRow = (5/3)*3 = 1*3 = 3
	// startCol = (7/3)*3 = 2*3 = 6
	// So we check the box starting at (3,6)

	// Now check all 9 spots in that 3×3 box
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if grid[i][j] == num {
				// Found the number in this box!
				return false // Not allowed! 🚫
			}
		}
	}

	// STEP 3: All checks passed!
	return true // Safe to place the number! ✅
}

// ============================================
// TOOL #2: THE PUZZLE SOLVER
// Job: Fill in all the empty spots
// ============================================
func SolveSudoku(grid *[9][9]int) bool {
	// Returns: true (solved!) or false (impossible!)

	// STEP 1: Look through EVERY cell in the grid
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			// STEP 2: Found an empty cell? (0 means empty)
			if grid[row][col] == 0 {

				// STEP 3: Try numbers 1, 2, 3, 4, 5, 6, 7, 8, 9
				for num := 1; num <= 9; num++ {

					// STEP 4: Ask the safety checker
					if isValid(grid, row, col, num) {
						// This number is safe! Place it!
						grid[row][col] = num

						// STEP 5: Try to solve the REST of the puzzle
						// 🤯 MAGIC: The function calls ITSELF!
						if SolveSudoku(grid) {
							// Success! Everything worked!
							return true // Keep this number ✅
						}

						// STEP 6: BACKTRACK!
						// That didn't work. Erase and try next number
						grid[row][col] = 0 // Put it back to empty
					}
				}

				// STEP 7: None of 1-9 worked here
				// Tell the previous step to try something else
				return false // Can't solve from here 🚫
			}
		}
	}

	// STEP 8: No empty cells left!
	return true // Puzzle is completely solved! 🎉
}

// ============================================
// TOOL #3: THE PRINTER
// Job: Show the solved puzzle on screen
// ============================================
func PrintSudoku(grid [9][9]int) {
	// Go through each row
	for i := 0; i < 9; i++ {
		// Go through each column in this row
		for j := 0; j < 9; j++ {
			// Print the number
			fmt.Print(grid[i][j])
			// Add space between numbers (but not after the last one)
			if j != 8 {
				fmt.Print(" ")
			}
		}
		// Move to next line after each row
		fmt.Println()
	}
}
