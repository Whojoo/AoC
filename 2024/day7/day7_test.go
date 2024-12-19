package day7_test

import (
	"testing"

	"github.com/Whojoo/AoC/2024/day7"
)

func TestCalculateTotalCalibrationResult(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input  []string
		result uint64
	}{
		"Given example input": {
			input: []string{
				"190: 10 19",
				"3267: 81 40 27",
				"83: 17 5",
				"156: 15 6",
				"7290: 6 8 6 15",
				"161011: 16 10 13",
				"192: 17 8 14",
				"21037: 9 7 18 13",
				"292: 11 6 16 20",
			},
			result: 3749,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mappedInput := day7.MapInput(test.input)
			got := day7.CalculateTotalCalibrationResult(mappedInput)
			if got != test.result {
				t.Errorf("got %d, want %d", got, test.result)
			}
		})
	}
}
