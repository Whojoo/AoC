package day12

import (
	"fmt"
	"iter"
)

type Assignment struct{}

func NewAssignment() *Assignment { return new(Assignment) }

func (a Assignment) FileName() string { return "day12.txt" }

func (a Assignment) Part1(input []string) int {
	garden := createPrimitiveGarden(input)

	price := 0
	visited := NewPlotSet()

	for row, gardenPlots := range garden {
		for column := range gardenPlots {
			if visited.Contains(row, column) {
				continue
			}

			price += CalculateAreaPerimeterPrice(row, column, garden, visited)
		}
	}

	return price
}

func (a Assignment) Part2(input []string) int {
	garden := createPrimitiveGarden(input)
	garden.RenderGarden()
	return len(input)
}

func CalculateAreaPerimeterPrice(row, column int, garden Garden, visited *PlotSet) int {
	const maxPerimeter = 4
	queue := NewPlotQueue()
	queue.Queue(Plot{row, column, garden[row][column]})

	var perimeter, area int

	for !queue.IsEmpty() {
		currentPlot := queue.Dequeue()

		if visited.Contains(currentPlot.row, currentPlot.column) {
			continue
		}

		visited.Add(currentPlot.row, currentPlot.column)

		currentPerimeter := maxPerimeter
		neighbours := garden.GetNeighbours(currentPlot.row, currentPlot.column)

		for neighbour := range neighbours {
			if neighbour.plotType == garden[row][column] {
				currentPerimeter--

				if !visited.Contains(neighbour.row, neighbour.column) {
					queue.Queue(neighbour)
				}
			}
		}

		area++
		perimeter += currentPerimeter
	}

	return perimeter * area
}

type PlotQueue []Plot

func NewPlotQueue() *PlotQueue { return new(PlotQueue) }

func (q *PlotQueue) Len() int        { return len(*q) }
func (q *PlotQueue) Peek() Plot      { return (*q)[0] }
func (q *PlotQueue) IsEmpty() bool   { return len(*q) == 0 }
func (q *PlotQueue) Queue(plot Plot) { *q = append(*q, plot) }

func (q *PlotQueue) Dequeue() Plot {
	if len(*q) == 0 {
		panic("Empty queue")
	}

	oldQueue := *q
	plot := oldQueue[0]
	*q = oldQueue[1:]
	return plot
}

type Plot struct {
	row, column int
	plotType    rune
}

type PlotSet map[struct{ row, column int }]struct{}

func NewPlotSet() *PlotSet { return &PlotSet{} }

func (ps *PlotSet) Add(row, column int) bool {
	plot := struct{ row, column int }{row, column}

	if ps.containsInternal(plot) {
		return false
	}

	(*ps)[plot] = struct{}{}
	return true
}

func (ps *PlotSet) Contains(row, column int) bool {
	return ps.containsInternal(struct{ row, column int }{row, column})
}

func (ps *PlotSet) containsInternal(plot struct{ row, column int }) bool {
	_, found := (*ps)[plot]
	return found
}

type Garden [][]rune

func (g Garden) Get(row, column int) (value rune, found bool) {
	if row < 0 || column < 0 || row >= len(g) || column >= len(g[row]) {
		return 0, false
	}

	return g[row][column], true
}

func (g Garden) GetNeighbours(row, column int) iter.Seq[Plot] {
	return func(yield func(Plot) bool) {
		if plotType, ok := g.Get(row+1, column); ok {
			plot := Plot{row + 1, column, plotType}
			if !yield(plot) {
				return
			}
		}

		if plotType, ok := g.Get(row-1, column); ok {
			plot := Plot{row - 1, column, plotType}
			if !yield(plot) {
				return
			}
		}

		if plotType, ok := g.Get(row, column-1); ok {
			plot := Plot{row, column - 1, plotType}
			if !yield(plot) {
				return
			}
		}

		if plotType, ok := g.Get(row, column+1); ok {
			plot := Plot{row, column + 1, plotType}
			if !yield(plot) {
				return
			}
		}
	}
}

func (g Garden) RenderGarden() {
	width := len(g[0])*2 + 1

	var emptyRow string
	for range width {
		emptyRow += "."
	}

	fmt.Println(emptyRow)

	for _, gardenRow := range g {
		renderRow := "."

		for _, plot := range gardenRow {
			renderRow += string(plot) + "."
		}

		fmt.Println(renderRow)
	}

	fmt.Println(emptyRow)
}

func createPrimitiveGarden(input []string) Garden {
	garden := make(Garden, len(input))

	for row, line := range input {
		garden[row] = make([]rune, len(line))
		for column, letterRune := range line {
			garden[row][column] = letterRune
		}
	}

	return garden
}
