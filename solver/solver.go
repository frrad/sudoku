package solver

import (
	"fmt"

	"github.com/frrad/boolform/bf"
	"github.com/frrad/boolform/bfgosat"
)

func Solve(specify [][3]int) map[[2]int]int {
	prob := basicProblem()

	for _, spec := range specify {
		prob.Assert(bf.Eq(bVar(spec[0], spec[1], spec[2]-1), bf.True))
	}

	sol := bfgosat.Solve(prob.AsFormula())

	if sol == nil {
		panic("no sol")
	}
	return decode(sol)
}

func basicProblem() *bf.Problem {
	prob := bf.NewProb()

	// Each number must be well defined
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			acc := []bf.Formula{}
			for k := 0; k < 9; k++ {
				acc = append(acc, bVar(i, j, k))
			}
			prob.Assert(bf.Unique(acc...))
		}
	}

	// Each row (col) must contain each number exactly once
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			acc1 := []bf.Formula{}
			acc2 := []bf.Formula{}
			for k := 0; k < 9; k++ {
				acc1 = append(acc1, bVar(i, k, j))
				acc2 = append(acc2, bVar(k, i, j))
			}
			prob.Assert(bf.Unique(acc1...))
			prob.Assert(bf.Unique(acc2...))
		}
	}

	// Each square has each number exactly once
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {

			for k := 0; k < 9; k++ {
				acc := []bf.Formula{}
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						acc = append(acc, bVar(i+3*a, j+3*b, k))
					}
				}
				prob.Assert(bf.Unique(acc...))
			}
		}
	}
	return prob
}

func decode(sol map[string]bool) map[[2]int]int {
	retval := map[[2]int]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			ans := -1
			for k := 0; k < 9; k++ {
				if sol[bVar(i, j, k).String()] {
					if ans != -1 {
						panic("should not happen")
					}
					ans = k
				}
			}
			retval[[2]int{i, j}] = ans + 1

		}
	}
	return retval
}

func bVar(x, y, z int) bf.Formula {
	return bf.Var(fmt.Sprintf("z-%d-%d-%d", x, y, z))
}
