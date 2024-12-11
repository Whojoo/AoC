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
	assignment := day1.CreateAssignment(input)
	want := 11

	// Act
	got := assignment.HandleFirst()

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
	assignment := day1.CreateAssignment(input)
	want := 31

	// Act
	got := assignment.HandleSecond()

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
