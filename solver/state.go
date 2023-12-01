package solver

import (
	"fmt"

	"github.com/mortawe/colour-bottles-solver/bottle"
)

type State struct {
	Bottles bottle.BottlePool
	Steps   Steps
}

func (s State) Copy() State {
	return State{
		Bottles: s.Bottles.Copy(),
		Steps:   s.Steps.Copy(),
	}
}

type Step struct {
	From, To int
}

func (s Step) String() string {
	return fmt.Sprintf("%d -> %d", s.From, s.To)
}

type Steps []Step

func (s Steps) Copy() Steps {
	sC := make(Steps, len(s))
	copy(sC, s)
	return sC
}

func (s State) getChildStates() []State {
	var childStates []State
	for idxFrom, from := range s.Bottles {
		if from.Filled() || from.IsEmpty() {
			continue
		}
		for idxTo, to := range s.Bottles {
			if idxTo == idxFrom || !bottle.CheckMove(from, to) {
				continue
			}
			if _, count := from.GetTopCombination(); count > to.RemainingCapacity() {
				continue
			}
			stateCopy := s.copyAndMove(idxFrom, idxTo)
			childStates = append(childStates, stateCopy)
		}
	}
	return childStates
}

func (s State) copyAndMove(idxFrom, idxTo int) State {
	stateCopy := s.Copy()
	fromCopy, toCopy := stateCopy.Bottles[idxFrom], stateCopy.Bottles[idxTo]
	for bottle.CheckMove(fromCopy, toCopy) {
		_ = bottle.Move(&fromCopy, &toCopy)
		stateCopy.Bottles[idxFrom], stateCopy.Bottles[idxTo] = fromCopy, toCopy
		stateCopy.Steps = append(stateCopy.Steps, Step{
			From: idxFrom,
			To:   idxTo,
		})
	}
	return stateCopy
}
