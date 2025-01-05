package day2_test

import (
	"testing"

	"github.com/Whojoo/AoC/2024/day2"
)

func TestHandleFirst(t *testing.T) {
	t.Parallel()

	// Arrange
	input := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	want := 2

	// Act
	got := day2.GetAssignment().Part1(input)

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandleSecond(t *testing.T) {
	t.Parallel()

	// Arrange
	input := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	want := 4

	// Act
	got := day2.GetAssignment().Part2(input)

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
