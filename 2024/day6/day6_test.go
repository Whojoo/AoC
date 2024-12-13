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
	got := day6.HandleFirst(input)

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandleSecond(t *testing.T) {
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
	got := day6.HandleSecond(input)

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
