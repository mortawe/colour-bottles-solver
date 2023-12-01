package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mortawe/colour-bottles-solver/bottle"
	"github.com/mortawe/colour-bottles-solver/solver"
)

func main() {
	file := flag.String("file", "example.json", "file with task to solve")
	flag.Parse()

	startTime := time.Now()
	bottles, err := bottle.ParseFile(*file)
	if err != nil {
		log.Fatal(err)
	}
	solution := solver.Solve(bottles)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Duration: ", duration)
	if solution == nil {
		fmt.Println("No solution exists")
		return
	}
	fmt.Println("Result: ", solution.Bottles.String())
	fmt.Println("Steps number: ", len(solution.Steps))
	fmt.Println("Steps order:")
	for _, step := range solution.Steps {
		fmt.Println(step)
	}
}
