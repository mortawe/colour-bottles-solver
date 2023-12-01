package main_test

import (
	"testing"

	"github.com/mortawe/colour-bottles-solver/bottle"
	"github.com/mortawe/colour-bottles-solver/solver"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	var testFiles = []string{
		"tests/1.json",
		"tests/2.json",
	}

	for _, file := range testFiles {
		bottles, err := bottle.ParseFile(file)
		assert.Nil(t, err, "parse")
		res := solver.Solve(bottles)

		assertSolution(t, bottles, res.Steps)
	}
}

func assertSolution(t *testing.T, bottles bottle.BottlePool, steps solver.Steps) {
	for _, step := range steps {
		ok := bottle.CheckMove(bottles[step.From], bottles[step.To])
		if !assert.True(t, ok, "check move") {
			return
		}
		ok = bottle.Move(&bottles[step.From], &bottles[step.To])
		if !assert.True(t, ok, "move") {
			return
		}
	}
	if !assert.True(t, bottles.Check(), "check") {
		return
	}
}
