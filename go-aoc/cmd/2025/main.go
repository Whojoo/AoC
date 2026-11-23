package main

import (
	"fmt"
	"go-aoc/internal/aoc2025"
	"go-aoc/internal/input"
)

func main() {
	in, err := input.ReadFromCli()
	if err != nil {
		panic(err)
	}

	p1, p2, err := aoc2025.Day1(in)
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1: %d\npart 2: %d\n", p1, p2)
}
