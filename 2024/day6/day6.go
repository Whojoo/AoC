package day6

import (
	"fmt"
	"strings"
)

func HandleFirst(input []string) int {
	grid := createGrid(input)
	// Generate gif https://github.com/JulianStremel/advent-of-code-24/blob/main/day6/day6.go

	x, y := grid.playerX, grid.playerY

	for grid.MovePlayer() {
	}

	grid.field[y][x] = playerMarker

	for _, row := range grid.field {
		fmt.Println(row)
	}

	return grid.marked
}

func HandleSecond(input []string) int {
	return 0
}

const (
	playerMarker string = "^"
	objectMarker string = "#"
	walkedMarker string = "X"
)

func createGrid(input []string) Grid {
	field := make([][]string, len(input))
	playerX, playerY := 0, 0

	for y, line := range input {
		field[y] = strings.Split(line, "")

		for x, char := range field[y] {
			if char == playerMarker {
				playerX = x
				playerY = y
			}
		}
	}

	grid := Grid{
		field:     field,
		playerX:   playerX,
		playerY:   playerY,
		direction: up,
	}

	return grid
}

type Direction struct {
	X, Y int
}

var (
	up    = Direction{X: 0, Y: -1}
	right = Direction{X: 1, Y: 0}
	left  = Direction{X: -1, Y: 0}
	down  = Direction{X: 0, Y: 1}
)

func (d Direction) Turn() Direction {
	switch d {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	default:
		panic(fmt.Sprintf("Direction is screwd, started with %v", d))
	}
}

type Grid struct {
	field     [][]string
	playerX   int
	playerY   int
	direction Direction
	marked    int
}

func (g *Grid) SetupPlayerRotation() {
	for {
		newY := g.playerY + g.direction.Y
		newX := g.playerX + g.direction.X

		if newX < 0 || newY < 0 || newX >= len(g.field[0]) || newY >= len(g.field) {
			break
		}

		if g.field[newY][newX] == objectMarker {
			g.direction = g.direction.Turn()
			continue
		}

		break
	}
}

func (g *Grid) MovePlayer() bool {
	g.SetupPlayerRotation()
	g.playerX += g.direction.X
	g.playerY += g.direction.Y

	if g.playerX < 0 || g.playerY < 0 || g.playerX >= len(g.field[0]) || g.playerY >= len(g.field) {
		g.marked++
		return false
	}

	if g.field[g.playerY][g.playerX] != walkedMarker {
		g.marked++
		g.field[g.playerY][g.playerX] = walkedMarker
	}

	return true
}
