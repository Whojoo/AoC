package day15_test

import (
	"github.com/Whojoo/AoC/2024/day15"
	"github.com/Whojoo/AoC/2024/shared"
	"testing"
)

func TestAssignment_Part1(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		inputPath string
		want      int
	}{
		"Default example": {
			inputPath: "../exampleInput/day15-1.txt",
			want:      10092,
		},
		"Default example 2": {
			inputPath: "../exampleInput/day15-2.txt",
			want:      2028,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			input := shared.ReadInputWithWeirdTokenPrevention(test.inputPath)
			assignment := day15.Assignment{}

			// Act
			got := assignment.Part1(input)

			// Assert
			if got != test.want {
				t.Errorf("Assignment.Part1() = %d, want %d", got, test.want)
			}
		})
	}
}
