package solver

import (
	satsolver "github.com/frrad/boolform/bfgini"
	"github.com/frrad/boolform/smt"
)

func Solve(specify [][3]int) map[[2]int]int {
	prob := smt.NewProb()
	state := setupState(prob)
	basicProblem(prob, state)

	for _, spec := range specify {
		prob.Assert(state[spec[0]][spec[1]][spec[2]-1].Eq(prob.NewBoolConst(true)))
	}

	works := prob.Solve(satsolver.Solve)

	if works == false {
		panic("no sol")
	}

	return decode(state)
}

func setupState(p *smt.Problem) [][][]*smt.Bool {
	state := make([][][]*smt.Bool, 9)

	for i := 0; i < 9; i++ {
		state[i] = make([][]*smt.Bool, 9)
		for j := 0; j < 9; j++ {
			state[i][j] = make([]*smt.Bool, 9)
			for k := 0; k < 9; k++ {
				state[i][j][k] = p.NewBool()
			}
		}
	}
	return state
}

func basicProblem(prob *smt.Problem, state [][][]*smt.Bool) {

	// Each number must be well defined
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			acc := []*smt.Bool{}
			for k := 0; k < 9; k++ {
				acc = append(acc, state[i][j][k])
			}
			prob.Assert(acc[0].Unique(acc[1:]...))
		}
	}

	// Each row (col) must contain each number exactly once
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			acc1 := []*smt.Bool{}
			acc2 := []*smt.Bool{}
			for k := 0; k < 9; k++ {
				acc1 = append(acc1, state[i][k][j])
				acc2 = append(acc2, state[k][i][j])
			}
			prob.Assert(acc1[0].Unique(acc1[1:]...))
			prob.Assert(acc2[0].Unique(acc2[1:]...))
		}
	}

	// Each square has each number exactly once
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {

			for k := 0; k < 9; k++ {
				acc := []*smt.Bool{}
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						acc = append(acc, state[i+3*a][j+3*b][k])
					}
				}
				prob.Assert(acc[0].Unique(acc[1:]...))
			}
		}
	}

}

func decode(state [][][]*smt.Bool) map[[2]int]int {
	retval := map[[2]int]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			ans := -1
			for k := 0; k < 9; k++ {
				if state[i][j][k].SolVal() {
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
