package day10_test

import (
	"testing"

	"github.com/Whojoo/AoC/2024/day10"
)

func TestAssignment_Handle(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input   []string
		results []string
	}{
		"Default example": {
			input: []string{
				"89010123",
				"78121874",
				"87430965",
				"96549874",
				"45678903",
				"32019012",
				"01329801",
				"10456732",
			},
			results: []string{
				"36",
				"81",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assignment := day10.NewAssignment()
			c := make(chan string)
			go assignment.Handle(test.input, c)

			<-c
			got1 := <-c
			got2 := <-c
			<-c

			if got1 != test.results[0] {
				t.Errorf("got %v, want %v", got1, test.results[0])
			}
			if got2 != test.results[1] {
				t.Errorf("got %v, want %v", got2, test.results[1])
			}
		})
	}
}
