package day4_test

import (
	"testing"

	"github.com/Whojoo/AoC/2024/day4"
)

func TestHandleFirst(t *testing.T) {
	t.Parallel()

	// Arrange
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	want := 18

	// Act
	got := day4.GetAssignment().Part1(input)

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandleSecond(t *testing.T) {
	t.Parallel()

	// Arrange
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	want := 9

	// Act
	got := day4.GetAssignment().Part2(input)

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
