package day13_test

import (
	"testing"

	"github.com/Whojoo/AoC/2024/day13"
	"github.com/Whojoo/AoC/2024/shared"
)

func TestAssignment_Part1(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		inputPath string
		want      int
	}{
		"Default example": {
			inputPath: "../exampleInput/day13.txt",
			want:      480,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			input := shared.ReadInput(test.inputPath)
			assignment := day13.Assignment{}

			// Act
			got := assignment.Part1(input)

			// Assert
			if got != test.want {
				t.Errorf("Assignment.Part1() = %d, want %d", got, test.want)
			}
		})
	}
}
