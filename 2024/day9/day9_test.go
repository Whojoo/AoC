package day9_test

import (
	"testing"

	"github.com/Whojoo/AoC/2024/day9"
	"github.com/Whojoo/AoC/2024/shared"
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

			<-c
			got1 := <-c
			got2 := <-c
			<-c

			if got1 != test.results[0] {
				t.Errorf("got %s, want %s", got1, test.results[0])
			}

			if got2 != test.results[1] {
				t.Errorf("got %s, want %s", got2, test.results[1])
			}
		})
	}
}

func BenchmarkAssignment_Part2(b *testing.B) {
	input := shared.ReadInput("../input/day9.txt")
	assignment := day9.GetAssignment()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		assignment.Part2(input)
	}
}

func BenchmarkAssignment_Part2_Bonus(b *testing.B) {
	input := shared.ReadInput("../exampleInput/day9_bonus.txt")
	assignment := day9.GetAssignment()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		assignment.Part2(input)
	}
}

func BenchmarkAssignment_Part2_Evil(b *testing.B) {
	input := shared.ReadInput("../exampleInput/day9_evil.txt")
	assignment := day9.GetAssignment()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		assignment.Part2(input)
	}
}
