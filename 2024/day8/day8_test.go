package day8_test

import (
	"testing"

	"github.com/Whojoo/AoC/2024/day8"
)

func TestHandle(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input   []string
		results []string
	}{
		"Default example": {
			input: []string{
				"............",
				"........0...",
				".....0......",
				".......0....",
				"....0.......",
				"......A.....",
				"............",
				"............",
				"........A...",
				".........A..",
				"............",
				"............",
			},
			results: []string{
				"Day 8",
				"14",
				"34",
			},
		},
		"Should have 2": {
			input: []string{
				"..........",
				"..........",
				"..........",
				"....a.....",
				"..........",
				".....a....",
				"..........",
				"..........",
				"..........",
				"..........",
			},
			results: []string{
				"Day 8",
				"2",
				"5",
			},
		},
		"Should have 4": {
			input: []string{
				"..........",
				"..........",
				"..........",
				"....a.....",
				"........a.",
				".....a....",
				"..........",
				"..........",
				"..........",
				"..........",
			},
			results: []string{
				"Day 8",
				"4",
				"8",
			},
		},
		"Should ignore A": {
			input: []string{
				"..........",
				"..........",
				"..........",
				"....a.....",
				"........a.",
				".....a....",
				"..........",
				"......A...",
				"..........",
				"..........",
			},
			results: []string{
				"Day 8",
				"4",
				"8",
			},
		},
		"Example part 2": {
			input: []string{
				"T.........",
				"...T......",
				".T........",
				"..........",
				"..........",
				"..........",
				"..........",
				"..........",
				"..........",
				"..........",
			},
			results: []string{
				"Day 8",
				"3",
				"9",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			c := make(chan string)
			assignment := day8.GetAssignment()
			go assignment.Handle(test.input, c)

			<-c
			result := <-c
			result2 := <-c
			<-c

			if result != test.results[1] {
				t.Fatalf("want %s, got %s", test.results[1], result)
			}

			if result2 != test.results[2] {
				t.Fatalf("want %s, got %s", test.results[2], result)
			}
		})
	}
}

func TestAntiNodeSet_GetSetValue(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input  day8.Vector
		result int
	}{
		"1010": {
			input:  day8.Vector{X: 10, Y: 10},
			result: 1010,
		},
		"0101": {
			input:  day8.Vector{X: 1, Y: 1},
			result: 101,
		},
		"0000": {
			input:  day8.Vector{X: 0, Y: 0},
			result: 0,
		},
		"0550": {
			input:  day8.Vector{X: 50, Y: 5},
			result: 550,
		},
		"0505": {
			input:  day8.Vector{X: 5, Y: 5},
			result: 505,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			set := day8.NewAntiNodeSet()
			got := set.GetSetValue(test.input)

			if got != test.result {
				t.Fatalf("want %d, got %d", test.result, got)
			}
		})
	}
}
