package main

import (
	"fmt"

	"github.com/frrad/sudoku/solver"
)

func main() {
	sol := solver.Solve([][3]int{
		{0, 0, 1},
		{1, 1, 2},
		{2, 2, 3},
		{3, 3, 4},
		{4, 4, 5},
		{5, 5, 6},
		{6, 6, 7},
		{7, 7, 8},
		{8, 8, 9},
	})

	show(sol)
}

func show(ans map[[2]int]int) {
	hsep := "-------------------\n"

	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Print(hsep)
		}

		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				fmt.Print("\b|")
			}

			fmt.Printf("%d ", ans[[2]int{i, j}])
		}

		fmt.Print("\b|\n")
	}
	fmt.Print(hsep)

}
