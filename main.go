package main

import (
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/frrad/sudoku/solver"
)

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("WASM Go Initialized")

	registerCallback()
	<-c
}

func registerCallback() {
	js.Global().Set("solveSudoku", js.NewCallback(solveSudoku))
}

func solveSudoku([]js.Value) {
	userInput := getInput()
	ans := solver.Solve(userInput)
	writeOutput(ans)
}

func writeOutput(ans map[[2]int]int) {
	for k, v := range ans {
		setCell(k[0], k[1], v)
	}
}

func setCell(x, y, v int) {
	ansCellName := fmt.Sprintf("ans-%d-%d", x, y)
	js.Global().Get("document").Call("getElementById", ansCellName).Set("value", v)
}

func getInput() [][3]int {
	ans := [][3]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			thisCell := getCell(i, j)
			if thisCell >= 0 {
				ans = append(ans, [3]int{i, j, thisCell})
			}
		}
	}
	return ans
}

func getCell(x, y int) int {
	inputCellName := fmt.Sprintf("input-%d-%d", x, y)
	inString := js.Global().Get("document").Call("getElementById", inputCellName).Get("value").String()
	i, err := strconv.Atoi(inString)
	if err != nil {
		return -1
	}
	return i
}
