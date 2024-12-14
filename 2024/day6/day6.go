package day6

import (
	"fmt"
	"strings"
)

func HandleFirst(input []string) int {
	grid := createGrid(input)

	x, y := grid.playerX, grid.playerY

	for grid.MovePlayer() {
	}

	grid.field[y][x] = "X"

	for _, row := range grid.field {
		fmt.Println(row)
	}

	return grid.marked
}

func HandleSecond(input []string) int {
	grid := createGrid(input)
	x, y := grid.playerX, grid.playerY
	for grid.MovePlayer() {
	}

	objects := createObjectsCollection(grid)
	sum := 0

	for _, object := range objects {
		otherObjects := filter(objects, func(o Object) bool { return o.ID != object.ID })
		sum += CalculateLoopingGuardOptions(object, otherObjects, grid)
	}

	grid.field[y][x] = "X"
	for _, row := range grid.field {
		fmt.Println(row)
	}

	return sum
}

type Object struct {
	X, Y, ID int
}

func CalculateLoopingGuardOptions(current Object, others []Object, grid Grid) int {
	// 4 options
	// top left:     right & down: (X++, Y+1) && (X-1, Y++) => v (r.X-1, d.Y+1)
	// top right:    left & down:  (X--, Y-1) && (X-1, Y++) => < (l.X-1, d.Y-1)
	// bottom right: left && up:   (X--, Y-1) && (X+1, Y--) => ^ (l.X+1, u.Y-1)
	// bottom left:  right && up:  (X++, Y+1) && (X+1, Y--) => > (r.X+1, u.Y+1)
	topLeft, tlObjs := calculateTopLeft(current, others, grid)
	topRight, trObjs := calculateTopRight(current, others, grid)
	bottomRight, brObjs := calculateBottomRight(current, others, grid)
	bottomLeft, blObjs := calculateBottomLeft(current, others, grid)

	objects := make([]Object, 0)
	objects = append(objects, tlObjs...)
	objects = append(objects, trObjs...)
	objects = append(objects, brObjs...)
	objects = append(objects, blObjs...)

	for _, obj := range objects {
		grid.field[obj.Y][obj.X] = "O"
	}

	return topLeft + topRight + bottomRight + bottomLeft
}

func calculateTopLeft(current Object, others []Object, grid Grid) (int, []Object) {
	const lookDown string = "v"
	sum := 0
	potentialObjects := make([]Object, 0)

	rightOthers := filter(others, func(o Object) bool { return o.X > current.X && o.Y == current.Y+1 })
	downOthers := filter(others, func(o Object) bool { return o.X == current.X-1 && o.Y > current.Y })
	for _, r := range rightOthers {
		for _, d := range downOthers {
			if d.Y+1 >= len(grid.field) {
				continue
			}

			if r.X-1 < 0 {
				continue
			}

			if grid.field[d.Y+1][r.X-1] == lookDown {
				potentialObjects = append(potentialObjects, Object{X: r.X - 1, Y: d.Y + 1})
				sum++
			}
		}
	}

	return sum, potentialObjects
}

func calculateTopRight(current Object, others []Object, grid Grid) (int, []Object) {
	const lookLeft string = "<"
	sum := 0
	potentialObjects := make([]Object, 0)

	leftOthers := filter(others, func(o Object) bool { return o.X < current.X && o.Y == current.Y-1 })
	downOthers := filter(others, func(o Object) bool { return o.X == current.X-1 && o.Y > current.Y })
	for _, l := range leftOthers {
		for _, d := range downOthers {
			if d.Y-1 < 0 {
				continue
			}

			if l.X-1 < 0 {
				continue
			}

			if grid.field[d.Y-1][l.X-1] == lookLeft {
				potentialObjects = append(potentialObjects, Object{X: l.X - 1, Y: d.Y - 1})
				sum++
			}
		}
	}

	return sum, potentialObjects
}

func calculateBottomRight(current Object, others []Object, grid Grid) (int, []Object) {
	const lookUp string = "^"
	sum := 0
	potentialObjects := make([]Object, 0)

	leftOthers := filter(others, func(o Object) bool { return o.X < current.X && o.Y == current.Y-1 })
	upOthers := filter(others, func(o Object) bool { return o.X == current.X+1 && o.Y < current.Y })
	for _, l := range leftOthers {
		for _, u := range upOthers {
			if u.Y-1 < 0 {
				continue
			}

			if l.X+1 >= len(grid.field[0]) {
				continue
			}

			if grid.field[u.Y-1][l.X+1] == lookUp {
				potentialObjects = append(potentialObjects, Object{X: l.X + 1, Y: u.Y - 1})
				sum++
			}
		}
	}

	return sum, potentialObjects
}

func calculateBottomLeft(current Object, others []Object, grid Grid) (int, []Object) {
	const lookRight string = ">"
	sum := 0
	potentialObjects := make([]Object, 0)

	rightOthers := filter(others, func(o Object) bool { return o.X > current.X && o.Y == current.Y+1 })
	upOthers := filter(others, func(o Object) bool { return o.X == current.X+1 && o.Y < current.Y })
	for _, r := range rightOthers {
		for _, u := range upOthers {
			if u.Y+1 >= len(grid.field) {
				continue
			}

			if r.X+1 >= len(grid.field[0]) {
				continue
			}

			if grid.field[u.Y+1][r.X+1] == lookRight {
				potentialObjects = append(potentialObjects, Object{X: r.X + 1, Y: u.Y + 1})
				sum++
			}
		}
	}

	return sum, potentialObjects
}

func filter(objs []Object, f func(Object) bool) []Object {
	var result []Object
	for _, obj := range objs {
		if f(obj) {
			result = append(result, obj)
		}
	}

	return result
}

func createObjectsCollection(grid Grid) []Object {
	objects := make([]Object, 0)
	i := 0
	for y, row := range grid.field {
		for x, char := range row {
			if char == objectMarker {
				objects = append(objects, Object{X: x, Y: y, ID: i})
				i++
			}
		}
	}

	return objects
}

const (
	playerMarker string = "^"
	objectMarker string = "#"
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

	if !isWalkedMarker(g.field[g.playerY][g.playerX]) {
		g.marked++
	}

	// Overwrite rendering so the last usage is always shown (needed for part 2)
	g.field[g.playerY][g.playerX] = renderWalkedMarker(g.direction)

	return true
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

func isWalkedMarker(marker string) bool {
	const walkedMarkers string = "^>v<"
	return strings.Contains(walkedMarkers, marker)
}
