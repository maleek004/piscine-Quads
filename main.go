package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
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

	// Determine which 3×3 box we are in..
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

func printrow(startChar, midChar, endChar string, countMid int) string {
	if countMid < 0 {
		return startChar + "\n"
	}
	return startChar + strings.Repeat(midChar, countMid) + endChar + "\n"
}

func QuadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	countMid := x - 2
	var b strings.Builder

	b.WriteString(printrow("o", "-", "o", countMid))
	for i := 0; i < y-2; i++ {
		b.WriteString(printrow("|", " ", "|", countMid))
	}
	if y > 1 {
		b.WriteString(printrow("o", "-", "o", countMid))
	}
	return b.String()
}

func QuadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	countMid := x - 2
	var b strings.Builder

	b.WriteString(printrow("/", "*", "\\", countMid))
	for i := 0; i < y-2; i++ {
		b.WriteString(printrow("*", " ", "*", countMid))
	}
	if y > 1 {
		b.WriteString(printrow("\\", "*", "/", countMid))
	}
	return b.String()
}

func QuadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	countMid := x - 2
	var b strings.Builder

	b.WriteString(printrow("A", "B", "A", countMid))
	for i := 0; i < y-2; i++ {
		b.WriteString(printrow("B", " ", "B", countMid))
	}
	if y > 1 {
		b.WriteString(printrow("C", "B", "C", countMid))
	}
	return b.String()
}

func QuadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	countMid := x - 2
	var b strings.Builder

	b.WriteString(printrow("A", "B", "C", countMid))
	for i := 0; i < y-2; i++ {
		b.WriteString(printrow("B", " ", "B", countMid))
	}
	if y > 1 {
		b.WriteString(printrow("A", "B", "C", countMid))
	}
	return b.String()
}

func QuadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	countMid := x - 2
	var b strings.Builder

	b.WriteString(printrow("A", "B", "C", countMid))
	for i := 0; i < y-2; i++ {
		b.WriteString(printrow("B", " ", "B", countMid))
	}
	if y > 1 {
		b.WriteString(printrow("C", "B", "A", countMid))
	}
	return b.String()
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
	// read stdin
	reader := bufio.NewReader(os.Stdin)
	input, _ := io.ReadAll(reader)
	content := string(input)

	if strings.TrimSpace(content) == "" {
		fmt.Println("Not a quad function")
		return
	}

	// remove ONLY one trailing newline
	clean := strings.TrimRight(content, "\n")

	lines := strings.Split(clean, "\n")
	y := len(lines)
	x := len(lines[0])

	// Expected regenerated quads
	target := clean + "\n"

	matches := []string{}

	if QuadA(x, y) == target {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", x, y))
	}
	if QuadB(x, y) == target {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", x, y))
	}
	if QuadC(x, y) == target {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", x, y))
	}
	if QuadD(x, y) == target {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", x, y))
	}
	if QuadE(x, y) == target {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", x, y))
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// alphabetical order
	sort.Strings(matches)

	fmt.Println(strings.Join(matches, " || "))
}
