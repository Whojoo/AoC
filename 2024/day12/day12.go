package day12

import (
	"strings"

	"github.com/Whojoo/AoC/2024/shared"
)

type Assignment struct{}

func NewAssignment() *Assignment { return new(Assignment) }

func (a Assignment) FileName() string { return "day12.txt" }

func (a Assignment) Part1(input []string) int {
	garden := createGarden(input)
	mappedGarden := make(map[int][]*GardenPlot)

	for _, gardenPlot := range garden {
		if gardenPlot.IsCreated {
			continue
		}

		mappedGarden[gardenPlot.ID] = gardenPlot.CreateGarden()
	}

	price := 0

	for _, gardenPlots := range mappedGarden {
		area := len(gardenPlots)
		perimeter := shared.Sum(gardenPlots, func(g *GardenPlot) int { return g.Perimeter() })
		plotPrice := area * perimeter
		price += plotPrice
	}

	return price
}

func (a Assignment) Part2(input []string) int {
	return len(input)
}

func createGarden(input []string) []*GardenPlot {
	currentID := 0
	garden := make([]*GardenPlot, 0)
	var previousGardenRow []*GardenPlot

	for _, gardenRow := range input {
		gardenPlotRow := shared.Project(strings.Split(gardenRow, ""), func(s string, _ int) *GardenPlot {
			gardenPlot := NewGardenPlot(currentID, s)
			currentID++
			return gardenPlot
		})

		for i, gardenPlot := range gardenPlotRow {
			if i > 0 {
				gardenPlot.AddNeighbour(gardenPlotRow[i-1])
				gardenPlotRow[i-1].AddNeighbour(gardenPlot)
			}

			if len(previousGardenRow) > 0 {
				gardenPlot.AddNeighbour(previousGardenRow[i])
				previousGardenRow[i].AddNeighbour(gardenPlot)
			}
		}

		previousGardenRow = gardenPlotRow
		garden = append(garden, gardenPlotRow...)
	}

	return garden
}

type GardenPlot struct {
	ID         int
	IsCreated  bool
	Neighbours []*GardenPlot
	PlantType  string
}

func NewGardenPlot(id int, plantType string) *GardenPlot {
	return &GardenPlot{
		ID:         id,
		IsCreated:  false,
		Neighbours: make([]*GardenPlot, 0),
		PlantType:  plantType,
	}
}

func (g *GardenPlot) AddNeighbour(neighbour *GardenPlot) {
	g.Neighbours = append(g.Neighbours, neighbour)
}

func (g *GardenPlot) Perimeter() int {
	const startingPerimeter = 4
	fellowPlots := shared.Filter(g.Neighbours, func(n *GardenPlot) bool { return g.ID == n.ID })
	return startingPerimeter - len(fellowPlots)
}

func (g *GardenPlot) ConnectToGarden(id int) []*GardenPlot {
	g.ID = id
	g.IsCreated = true
	fellowPlots := []*GardenPlot{g}

	for _, neighbour := range g.Neighbours {
		if neighbour.ID != id && neighbour.PlantType == g.PlantType {
			fellowPlots = append(fellowPlots, neighbour.ConnectToGarden(id)...)
		}
	}

	return fellowPlots
}

func (g *GardenPlot) CreateGarden() []*GardenPlot {
	fellowPlots := []*GardenPlot{g}
	g.IsCreated = true

	for _, neighbour := range g.Neighbours {
		if neighbour.PlantType == g.PlantType && neighbour.ID != g.ID {
			fellowPlots = append(fellowPlots, neighbour.ConnectToGarden(g.ID)...)
		}
	}

	return fellowPlots
}
