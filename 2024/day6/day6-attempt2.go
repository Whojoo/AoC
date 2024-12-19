package day6

import (
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/Whojoo/AoC/2024/shared"
)

type AssignmentAttempt2 struct{}

func GetAssignment2() AssignmentAttempt2 {
	return AssignmentAttempt2{}
}

func (AssignmentAttempt2) Handle(input []string, c chan<- string) {
	startTime := time.Now()
	grid := createGrid(input)
	// Need a copy for the looping later
	gridCopy := grid.Copy()

	x, y := grid.playerX, grid.playerY

	for grid.RunGuardRoute() {
	}

	first := grid.CountMarkedTiles()
	firstTime := time.Since(startTime)

	startTime = time.Now()

	tilesToCheck := shared.Filter(grid.GetWalkedTiles(), func(t Tile) bool { return !(t.X == x && t.Y == y) })
	guardLoopsIncrementer := GuardLoopsIncrementer{}
	var wg sync.WaitGroup

	for _, tile := range tilesToCheck {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Need a new copy since all goroutines want to fiddle with grid
			checkingGrid := gridCopy.Copy()
			checkingGrid.MarkTileAsObject(tile)
			for checkingGrid.RunGuardRouteWithLoopCheck(&guardLoopsIncrementer) {
			}
		}()
	}

	wg.Wait()

	second := guardLoopsIncrementer.Get()
	secondTime := time.Since(startTime)

	c <- "Day 6"
	c <- fmt.Sprintf("First result: %d in %s", first, firstTime)
	c <- fmt.Sprintf("Second result: %d in %s", second, secondTime)

	close(c)
}

func (AssignmentAttempt2) FileName() string {
	return "day6.txt"
}

func (g *Grid) RunGuardRoute() bool {
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

	return true
}

func (g *Grid) RunGuardRouteWithLoopCheck(loopIncrementer *GuardLoopsIncrementer) bool {
	g.SetupPlayerRotation()
	g.playerX += g.direction.X
	g.playerY += g.direction.Y

	if !g.IsInsideGrid(g.playerX, g.playerY) {
		return false
	}

	tile := &g.field[g.playerY][g.playerX]

	if tile.State == walked && slices.Contains(tile.WalkedDirections, g.direction) {
		loopIncrementer.Increment()
		return false
	}

	tile.State = walked
	tile.WalkedDirections = append(tile.WalkedDirections, g.direction)
	tile.Render = renderWalkedMarker(g.direction)

	return true
}

func (g *Grid) MarkTileAsObject(t Tile) {
	tile := &g.field[t.Y][t.X]
	tile.State = obstruction
	tile.Render = objectMarker
}

func (g *Grid) GetWalkedTiles() []Tile {
	walkedTiles := make([]Tile, 0)

	for _, row := range g.field {
		walkedTiles = append(walkedTiles, shared.Filter(row, func(t Tile) bool { return t.State == walked })...)
	}

	return walkedTiles
}

type GuardLoopsIncrementer struct {
	loops int
	mu    sync.Mutex
}

func (g *GuardLoopsIncrementer) Increment() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.loops++
}

func (g *GuardLoopsIncrementer) Get() int {
	return g.loops
}
