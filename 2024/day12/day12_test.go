package day12_test

import (
	"testing"

	"github.com/Whojoo/AoC/2024/day12"
	"github.com/Whojoo/AoC/2024/shared"
)

func TestAssignment_Part1(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		inputPath string
		want      int
	}{
		"Default Example": {
			inputPath: "../exampleInput/day12.txt",
			want:      1930,
		},
		"Extra Example 1": {
			inputPath: "../exampleInput/day12_1.txt",
			want:      140,
		},
		"Extra Example 2": {
			inputPath: "../exampleInput/day12_2.txt",
			want:      772,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			input := shared.ReadInput(test.inputPath)
			assignment := day12.NewAssignment()

			// Act
			got := assignment.Part1(input)

			// Assert
			if got != test.want {
				t.Errorf("Assignment.Part1() = %d, want %d", got, test.want)
			}
		})
	}
}

func TestAssignment_Part2(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		inputPath string
		want      int
	}{
		"Default Example": {
			inputPath: "../exampleInput/day12.txt",
			want:      1206,
		},
		"Extra Example 1": {
			inputPath: "../exampleInput/day12_1.txt",
			want:      80,
		},
		"Extra Example 2": {
			inputPath: "../exampleInput/day12_2.txt",
			want:      436,
		},
		"E made out of E's": {
			inputPath: "../exampleInput/day12_3.txt",
			want:      236,
		},
		"ABBA": {
			inputPath: "../exampleInput/day12_4.txt",
			want:      368,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			input := shared.ReadInput(test.inputPath)
			assignment := day12.NewAssignment()

			// Act
			got := assignment.Part2(input)

			// Assert
			if got != test.want {
				t.Errorf("Assignment.Part1() = %d, want %d", got, test.want)
			}
		})
	}
}

var global int

func BenchmarkAssignment_Part1(b *testing.B) {
	var local int
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := shared.ReadInput("../input/day12.txt")
		assignment := day12.NewAssignment()
		b.StartTimer()

		local = assignment.Part1(input)
	}

	global = local
}
