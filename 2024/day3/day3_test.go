package day3_test

import (
	"testing"

	"github.com/Whojoo/AoC/2024/day3"
)

func TestHandleFirst(t *testing.T) {
	// Arrange
	input := []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}
	want := 161

	// Act
	got := day3.HandleFirst(input)

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandleSecond(t *testing.T) {
	// Arrange
	input := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}
	want := 48

	// Act
	got := day3.HandleSecond(input)

	// Assert
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
