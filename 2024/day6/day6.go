package day6

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Whojoo/AoC/2024/shared"
)

type Assignment struct{}

func GetAssignment() Assignment {
	return Assignment{}
}

func (Assignment) Handle(input []string, c chan<- int) {
	grid := createGrid(input)

	x, y := grid.playerX, grid.playerY

	for grid.MovePlayer() {
	}

	grid.field[y][x].Render = "X"

	c <- grid.CountMarkedTiles()
	c <- grid.CountPotentialLoopObjects()

	close(c)
}

func (Assignment) FileName() string {
	return "day6.txt"
}

type Object struct {
	X, Y, ID int
}

const (
	playerMarker string = "^"
	objectMarker string = "#"
)

func createGrid(input []string) Grid {
	tilesField := make([][]Tile, len(input))
	playerX, playerY := 0, 0

	for y, line := range input {
		fields := strings.Split(line, "")
		tilesField[y] = make([]Tile, len(fields))

		for x, field := range fields {
			tilesField[y][x] = Tile{
				State:            getTileState(field),
				Render:           field,
				X:                x,
				Y:                y,
				WalkedDirections: []Direction{},
			}

			if tilesField[y][x].State == walked {
				playerX = x
				playerY = y
				tilesField[y][x].WalkedDirections = append(tilesField[y][x].WalkedDirections, up)
			}
		}
	}

	grid := Grid{
		field:     tilesField,
		playerX:   playerX,
		playerY:   playerY,
		direction: up,
	}

	return grid
}

func getTileState(s string) Marker {
	switch s {
	case playerMarker:
		return walked
	case objectMarker:
		return obstruction
	default:
		return empty
	}
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

type Marker int

const (
	empty Marker = iota
	walked
	obstruction
)

type Tile struct {
	State                 Marker
	Render                string
	X, Y                  int
	WalkedDirections      []Direction
	CanBeUsedAsLoopObject bool
}

func (t *Tile) Copy() Tile {
	newTile := Tile{
		State:                 t.State,
		Render:                t.Render,
		X:                     t.X,
		Y:                     t.Y,
		CanBeUsedAsLoopObject: t.CanBeUsedAsLoopObject,
	}
	newTile.WalkedDirections = append(newTile.WalkedDirections, t.WalkedDirections...)

	return newTile
}

type Grid struct {
	field     [][]Tile
	playerX   int
	playerY   int
	direction Direction
}

func (g *Grid) SetupPlayerRotation() {
	for {
		newY := g.playerY + g.direction.Y
		newX := g.playerX + g.direction.X

		if !g.IsInsideGrid(newX, newY) {
			break
		}

		if g.field[newY][newX].State == obstruction {
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

	if !g.IsInsideGrid(g.playerX, g.playerY) {
		return false
	}

	tile := &g.field[g.playerY][g.playerX]
	tile.State = walked
	tile.WalkedDirections = append(tile.WalkedDirections, g.direction)
	tile.Render = renderWalkedMarker(g.direction)

	gridCopy := g.Copy()
	checkForLoop(gridCopy, tile)

	return true
}

func (g *Grid) Copy() Grid {
	newGrid := Grid{
		field:     make([][]Tile, len(g.field)),
		playerX:   g.playerX,
		playerY:   g.playerY,
		direction: g.direction,
	}

	for y, row := range g.field {
		newGrid.field[y] = make([]Tile, len(row))
		for x, tile := range row {
			newGrid.field[y][x] = tile.Copy()
		}
	}

	return newGrid
}

func (g *Grid) IsInsideGrid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(g.field[0]) && y < len(g.field)
}

func checkForLoop(g Grid, guardTile *Tile) {
	oX, oY := guardTile.X+g.direction.X, guardTile.Y+g.direction.Y

	// Early out if next step is outside the map
	if !g.IsInsideGrid(oX, oY) {
		return
	}

	newObjectTile := &g.field[oY][oX]
	if newObjectTile.State == obstruction {
		return
	}
	if len(newObjectTile.WalkedDirections) > 0 {
		// Early out if the object is placed on a path we visited before, otherwise we wouldn't have got here
		return
	}

	newObjectTile.State = obstruction
	newObjectTile.Render = "O"

	for {
		// Walk the guard
		g.SetupPlayerRotation()
		g.playerX += g.direction.X
		g.playerY += g.direction.Y

		if !g.IsInsideGrid(g.playerX, g.playerY) {
			return
		}

		tile := &g.field[g.playerY][g.playerX]

		if tile.State == walked && slices.Contains(tile.WalkedDirections, g.direction) {
			guardTile.CanBeUsedAsLoopObject = true
			// g.Print()
			return
		}

		tile.State = walked
		tile.WalkedDirections = append(tile.WalkedDirections, g.direction)
		tile.Render = renderWalkedMarker(g.direction)
	}
}

func (g *Grid) CountMarkedTiles() int {
	sum := 0

	for _, row := range g.field {
		walkedTiles := shared.Filter(row, func(t Tile) bool { return t.State == walked })
		sum += len(walkedTiles)
	}

	return sum
}

func (g *Grid) CountPotentialLoopObjects() int {
	sum := 0

	for _, row := range g.field {
		potentialLoopObjects := shared.Filter(row, func(t Tile) bool { return t.CanBeUsedAsLoopObject })
		sum += len(potentialLoopObjects)
	}

	return sum
}

func (g *Grid) Print() {
	for _, field := range g.field {
		fieldRenders := shared.Project(field, func(t Tile) string { return t.Render })
		fmt.Println(strings.Join(fieldRenders, " "))
	}
	fmt.Println()
}

func renderWalkedMarker(d Direction) string {
	switch d {
	case up:
		return "^"
	case right:
		return ">"
	case left:
		return "<"
	case down:
		return "v"
	default:
		panic(fmt.Sprintf("Direction is screwd, started with %v", d))
	}
}
