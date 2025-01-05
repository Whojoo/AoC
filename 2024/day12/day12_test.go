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
		"Extra Example 3": {
			inputPath: "../exampleInput/day12_3.txt",
			want:      236,
		},
		"Extra Example 4": {
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
