package day9_test

import (
	"github.com/Whojoo/AoC/2024/day9"
	"testing"
)

func TestAssignment_Handle(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input   []string
		results []string
	}{
		"Default example": {
			input: []string{"2333133121414131402"},
			results: []string{
				"1928",
				"2858",
			},
		},
		"12345": {
			input: []string{"12345"},
			results: []string{
				"60",
				"132",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			c := make(chan string)
			assignment := day9.GetAssignment()
			go assignment.Handle(test.input, c)

			_ = <-c
			got1 := <-c
			got2 := <-c
			_ = <-c

			if got1 != test.results[0] {
				t.Errorf("got %s, want %s", got1, test.results[0])
			}

			if got2 != test.results[1] {
				t.Errorf("got %s, want %s", got2, test.results[1])
			}
		})
	}
}
