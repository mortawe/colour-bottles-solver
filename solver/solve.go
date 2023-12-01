package solver

import (
	"github.com/mortawe/colour-bottles-solver/bottle"
	"github.com/mortawe/colour-bottles-solver/stack"
)

func Solve(bottles []bottle.Bottle) *State {
	visitedState := make(map[string]struct{})

	var searchOrder stack.Stack[State]

	searchOrder.Push(State{
		Bottles: bottles,
	})
	for !searchOrder.IsEmpty() {
		state, _ := searchOrder.Pop()
		hash := state.Bottles.StringSorted()
		if _, ok := visitedState[hash]; ok {
			continue
		}
		visitedState[hash] = struct{}{}
		if state.Bottles.Check() {
			return &state
		}
		childStates := state.getChildStates()
		for _, s := range childStates {
			searchOrder.Push(s)
		}
	}
	return nil
}
