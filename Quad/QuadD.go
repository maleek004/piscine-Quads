package Quad

import (
	"fmt"
	"strings"
)

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	countMidCells := x - 2

	printrow := func(startChar string, midChar string, endChar string) {
		fmt.Print(startChar) //
		if countMidCells >= 0 {
			fmt.Print(strings.Repeat(midChar, countMidCells))
			fmt.Println(endChar)
		} else {
			fmt.Println()
		}

	}
	//  Now using the helper function
	printrow("A", "B", "C")

	// if we have any middle row i.e y >2
	for i := y - 2; i >= 1; i-- {
		printrow("B", " ", "B")
	}
	if y > 1 {
		printrow("A", "B", "C")
	}

}
