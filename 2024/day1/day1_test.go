package day1_test

import (
	"testing"

	"github.com/Whojoo/AoC/2024/day1"
)

func TestHandleFirst(t *testing.T) {
	t.Parallel()

	// Arrange
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	want := 11

	// Act
	got := day1.GetAssignment().Part1(input)

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandleSecond(t *testing.T) {
	t.Parallel()

	// Arrange
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	want := 31

	// Act
	got := day1.GetAssignment().Part2(input)

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
