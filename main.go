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
	userInput, err := getInput()
	if err != nil {
		updateMessage(fmt.Sprintf("invalid input: %+v", err))
		return
	}
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

func getInput() ([][3]int, error) {
	ans := [][3]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cellVal, isFilled, err := getCell(i, j)
			if err != nil {
				return nil, err
			}
			if isFilled {
				ans = append(ans, [3]int{i, j, cellVal})
			}
		}
	}
	return ans, nil
}

func getCell(x, y int) (int, bool, error) {
	inputCellName := fmt.Sprintf("input-%d-%d", x, y)
	inString := js.Global().Get("document").Call("getElementById", inputCellName).Get("value").String()
	if inString == "" {
		return 0, false, nil
	}

	i, err := strconv.Atoi(inString)
	if err != nil {
		return 0, false, fmt.Errorf("can't parse %s", inString)
	}
	return i, true, nil
}
