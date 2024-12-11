package day4

import (
	"slices"
	"strings"
)

func HandleFirst(input []string) int {
	// grid[y][x]
	grid := toGrid(input)
	sum := 0

	// Find X
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == "X" {
				sum += searchGrid(grid, x, y)
			}
		}
	}

	return sum
}

func HandleSecond(input []string) int {
	// grid[y][x]
	grid := toGrid(input)
	sum := 0

	// Find X
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == "A" {
				sum += searchGridForCrossedMAS(grid, x, y)
			}
		}
	}

	return sum
}

var firstLettersInOrder = []string{"X", "M", "A", "S"}

func toGrid(input []string) [][]string {
	grid := make([][]string, len(input))

	for i := range grid {
		grid[i] = strings.Split(input[i], "")
	}

	return grid
}

type direction struct {
	Y int
	X int
}

var (
	upLeft    = direction{-1, -1}
	up        = direction{-1, 0}
	upRight   = direction{-1, 1}
	left      = direction{0, -1}
	right     = direction{0, 1}
	downLeft  = direction{1, -1}
	down      = direction{1, 0}
	downRight = direction{1, 1}
)

func searchGrid(grid [][]string, x int, y int) int {
	return crawlGrid(grid, x, y, 1, upLeft, firstLettersInOrder) +
		crawlGrid(grid, x, y, 1, upRight, firstLettersInOrder) +
		crawlGrid(grid, x, y, 1, up, firstLettersInOrder) +
		crawlGrid(grid, x, y, 1, left, firstLettersInOrder) +
		crawlGrid(grid, x, y, 1, right, firstLettersInOrder) +
		crawlGrid(grid, x, y, 1, downLeft, firstLettersInOrder) +
		crawlGrid(grid, x, y, 1, down, firstLettersInOrder) +
		crawlGrid(grid, x, y, 1, downRight, firstLettersInOrder)
}

func searchGridForCrossedMAS(grid [][]string, x int, y int) int {
	cross1 := []string{"A", getFromGrid(grid, x, y, upLeft), getFromGrid(grid, x, y, downRight)}
	cross2 := []string{"A", getFromGrid(grid, x, y, upRight), getFromGrid(grid, x, y, downLeft)}

	slices.Sort(cross1)
	slices.Sort(cross2)
	expected := []string{"A", "M", "S"}

	if eq(cross1, expected) && eq(cross2, expected) {
		return 1
	}

	return 0
}

func eq(l []string, r []string) bool {
	if len(l) != len(r) {
		return false
	}

	for i := range l {
		if l[i] != r[i] {
			return false
		}
	}

	return true
}

func getFromGrid(grid [][]string, x int, y int, movementDirection direction) string {
	if !canCrawl(grid, x, y, movementDirection) {
		return ""
	}

	return grid[y+movementDirection.Y][x+movementDirection.X]
}

func crawlGrid(grid [][]string, x int, y int, depth int, movementDirection direction, word []string) int {
	if depth >= len(word) {
		return 0
	}

	if !canCrawl(grid, x, y, movementDirection) {
		return 0
	}

	newX, newY := x+movementDirection.X, y+movementDirection.Y

	if grid[newY][newX] != word[depth] {
		return 0
	}

	if grid[newY][newX] == word[len(word)-1] {
		return 1
	}

	return crawlGrid(grid, newX, newY, depth+1, movementDirection, word)
}

func canCrawl(grid [][]string, x int, y int, movementDirection direction) bool {
	newX, newY := x+movementDirection.X, y+movementDirection.Y

	isBelowZero := newX < 0 || newY < 0
	isLenOrMore := newY >= len(grid) || newX >= len(grid[y])

	return !isBelowZero && !isLenOrMore
}
