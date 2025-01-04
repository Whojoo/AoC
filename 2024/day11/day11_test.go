package day11_test

import (
	"github.com/Whojoo/AoC/2024/day11"
	"testing"
)

func TestAssignment_Handle(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input   []string
		results []string
	}{
		"Default example": {
			input: []string{
				"125 17",
			},
			results: []string{
				"55312",
				"0",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assignment := day11.NewAssignment()
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
