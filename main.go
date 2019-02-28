package main

import (
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/frrad/sudoku/solver"
)

func main() {
	c := make(chan struct{}, 0)
	updateMessage("solver loaded.")
	registerCallback()
	<-c
}

func registerCallback() {
	js.Global().Set("solveSudoku", js.NewCallback(solveSudoku))
}

func solveSudoku([]js.Value) {
	updateMessage("loading input...")
	userInput := getInput()
	updateMessage("solving...")
	ans, err := solver.Solve(userInput)
	if err != nil {
		updateMessage(fmt.Sprintf("failed to solve: %+v", err))
		return
	}
	updateMessage("solved!")
	bolds := map[[2]int]bool{}
	for _, input := range userInput {
		bolds[[2]int{input[0], input[1]}] = true
	}

	writeOutput(ans, bolds)
}

func updateMessage(msg string) {
	writeHTML("messageP", msg)
}

func writeOutput(ans map[[2]int]int, bolds map[[2]int]bool) {
	for k, v := range ans {
		x, y := k[0], k[1]
		setCell(x, y, v, bolds[[2]int{x, y}])
	}
}

func setCell(x, y, v int, bold bool) {
	ansCellName := fmt.Sprintf("ans-%d-%d", x, y)
	html := fmt.Sprintf("&nbsp;%d&nbsp;", v)
	if bold {
		html = fmt.Sprintf("<mark><b>%s</b></mark>", html)
	}
	writeHTML(ansCellName, html)
}

func writeHTML(id, html string) {
	js.Global().Get("document").Call("getElementById", id).Set("innerHTML", html)
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
