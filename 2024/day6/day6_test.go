package day6_test

import (
	"testing"

	"github.com/Whojoo/AoC/2024/day6"
)

func TestHandleFirst(t *testing.T) {
	// Arrange
	input := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}
	want := 41

	// Act
	c := make(chan int)
	go day6.GetAssignment().Handle(input, c)
	got := <-c
	// Skip second
	<-c

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

// Ignore function length linting for table driven test
//
//nolint:funlen
func TestHandleSecond_WhenGivenExampleInput_ShouldReturnCorrectResult(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name: "example 1",
			input: []string{
				"##..",
				"...#",
				"....",
				"^.#.",
			},
			want: 0,
		},
		{
			name: "example 2",
			input: []string{
				"..........",
				"..........",
				"..#.......",
				"........#.",
				"........#.",
				"..........",
				"....^.....",
				"..........",
				".#........",
				".......#..",
			},
			want: 1,
		},
		{
			name: "example 3",
			input: []string{
				"....#.....",
				".........#",
				"..........",
				"..#.......",
				".......#..",
				"..........",
				".#..^.....",
				"........#.",
				"#.........",
				"......#...",
			},
			want: 6,
		},
		{
			name: "example 4",
			input: []string{
				".#...",
				"....#",
				".....",
				".^.#.",
				"#....",
				"..#..",
			},
			want: 3,
		},
		{
			name: "example 5",
			input: []string{
				"....",
				"#...",
				".^#.",
				".#..",
			},
			want: 0,
		},
		{
			name: "example 6",
			input: []string{
				"##..",
				"...#",
				"....",
				"^.#.",
			},
			want: 0,
		},
	}

	for _, test := range tests {
		testName := test.name
		t.Run(testName, func(t *testing.T) {
			c := make(chan int)
			go day6.GetAssignment().Handle(test.input, c)
			<-c
			got := <-c

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}

func TestHandleSecond_WhenRunMultipleTimes_ShouldGiveTheSameResult(t *testing.T) {
	// Arrange
	input := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	const iterations int = 100
	channels := make([]chan int, iterations)

	// Act
	for i := 0; i < iterations; i++ {
		channels[i] = make(chan int)
		go func() {
			c := make(chan int)
			defer close(channels[i])

			go day6.GetAssignment().Handle(input, c)
			<-c
			result := <-c
			channels[i] <- result
		}()
	}

	results := make([]int, 0)
	for _, channel := range channels {
		for result := range channel {
			results = append(results, result)
		}
	}

	// Assert
	firstResult := results[0]
	for i := 1; i < iterations; i++ {
		if results[i] != firstResult {
			t.Errorf("results are not consistent")
		}
	}
}
