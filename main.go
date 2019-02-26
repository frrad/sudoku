package main

import (
	"fmt"
	"syscall/js"

	"github.com/frrad/sudoku/solver"
)

func main() {
	solveSudoku()
	fmt.Println("asdfk")
}

func solveSudoku() {
	ans := solver.Solve([][3]int{})
	for k, v := range ans {
		setCell(k[0], k[1], v)
	}
}

func setCell(x, y, v int) {
	ansCellName := fmt.Sprintf("ans-%d-%d", x, y)
	js.Global().Get("document").Call("getElementById", ansCellName).Set("value", v)
}
