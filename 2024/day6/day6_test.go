package day6_test

import (
	"github.com/Whojoo/AoC/2024/shared"
	"testing"

	"github.com/Whojoo/AoC/2024/day6"
)

func TestHandleFirst(t *testing.T) {
	t.Parallel()

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
	got := day6.GetAssignment().Part1(input)

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

// Ignore function length linting for table driven test
//
//nolint:funlen
func TestHandleSecond_WhenGivenExampleInput_ShouldReturnCorrectResult(t *testing.T) {
	t.Parallel()

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
		{
			name: "example 7",
			input: []string{
				"###.",
				"#.^.",
				"..#.",
			},
			want: 1,
		},
		{
			name: "example 8",
			input: []string{
				".#......",
				"........",
				"......#.",
				"........",
				"........",
				"...^....",
				"#.......",
				"..#..#..",
			},
			want: 1,
		},
	}

	for _, test := range tests {
		testName := test.name
		assignment := day6.GetAssignment()
		t.Run(testName, func(t *testing.T) {
			got := assignment.Part2(test.input)
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}

func BenchmarkAssignment_Part1(b *testing.B) {
	input := shared.ReadInput("../input/day6.txt")
	assignment := day6.GetAssignment()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		assignment.Part1(input)
	}
}

func BenchmarkAssignment_Part2(b *testing.B) {
	input := shared.ReadInput("../input/day6.txt")
	assignment := day6.GetAssignment()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		assignment.Part2(input)
	}
}
