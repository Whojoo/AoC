package day12

import (
	"iter"
	"slices"
)

type Assignment struct{}

func NewAssignment() *Assignment { return new(Assignment) }

func (a Assignment) FileName() string { return "day12.txt" }

func (a Assignment) Part1(input []string) int {
	garden := createGarden(input)

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
	garden := createGarden(input)
	// ignore this stuff garden.RenderGarden()

	price := 0
	visited := NewPlotSet()

	for row, gardenPlots := range garden {
		for column := range gardenPlots {
			if visited.Contains(row, column) {
				continue
			}

			price += CalculateAreaSidePrice(row, column, garden, visited)
		}
	}

	return price
}

func CalculateAreaSidePrice(row, column int, garden Garden, visited *PlotSet) int {
	queue := NewPlotQueue()
	queue.Queue(Plot{row, column, garden[row][column]})

	var area, sides int
	sideStore := make(SideStore)

	for !queue.IsEmpty() {
		currentPlot := queue.Dequeue()

		if visited.Contains(currentPlot.row, currentPlot.column) {
			continue
		}

		visited.Add(currentPlot.row, currentPlot.column)

		for plotSide := range garden.GetSides(currentPlot.row, currentPlot.column) {
			isSameType := plotSide.hasPlot && plotSide.plot.plotType == garden[row][column]
			isVisited := visited.Contains(plotSide.plot.row, plotSide.plot.column)

			if isSameType && !isVisited {
				queue.Queue(plotSide.plot)
			} else if !isSameType {
				sideStore.AddSide(plotSide.sideHeight, plotSide.sideLocation, plotSide.orientation, plotSide.sidePresence)
			}
		}

		area++
	}

	sides = 0
	for _, v := range sideStore {
		sides += len(v)
	}

	return sides * area
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

		for neighbour := range garden.GetNeighbours(currentPlot.row, currentPlot.column) {
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

type PlotSide struct {
	plot                     Plot
	hasPlot                  bool
	orientation              SideOrientation
	sidePresence             SidePresence
	sideHeight, sideLocation int
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

func (g Garden) GetSides(row, column int) iter.Seq[PlotSide] {
	return func(yield func(side PlotSide) bool) {
		plotType, ok := g.Get(row+1, column)
		plotSide := PlotSide{
			plot:         Plot{row + 1, column, plotType},
			hasPlot:      ok,
			orientation:  horizontal,
			sideHeight:   row + 1,
			sideLocation: column,
			sidePresence: left,
		}
		if !yield(plotSide) {
			return
		}

		plotType, ok = g.Get(row, column+1)
		plotSide = PlotSide{
			plot:         Plot{row, column + 1, plotType},
			hasPlot:      ok,
			orientation:  vertical,
			sideHeight:   column + 1,
			sideLocation: row,
			sidePresence: left,
		}
		if !yield(plotSide) {
			return
		}

		plotType, ok = g.Get(row-1, column)
		plotSide = PlotSide{
			plot:         Plot{row - 1, column, plotType},
			hasPlot:      ok,
			orientation:  horizontal,
			sideHeight:   row,
			sideLocation: column,
			sidePresence: right,
		}
		if !yield(plotSide) {
			return
		}

		plotType, ok = g.Get(row, column-1)
		plotSide = PlotSide{
			plot:         Plot{row, column - 1, plotType},
			hasPlot:      ok,
			orientation:  vertical,
			sideHeight:   column,
			sideLocation: row,
			sidePresence: right,
		}
		if !yield(plotSide) {
			return
		}
	}
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

func createGarden(input []string) Garden {
	garden := make(Garden, len(input))

	for row, line := range input {
		garden[row] = make([]rune, len(line))
		for column, letterRune := range line {
			garden[row][column] = letterRune
		}
	}

	return garden
}

type SideOrientation int

const (
	horizontal SideOrientation = iota
	vertical
)

type SidePresence int

const (
	left SidePresence = iota
	right
)

type Side struct {
	start, end   int
	orientation  SideOrientation
	sidePresence SidePresence
}

type SideStore map[int][]*Side

func (s *SideStore) AddSide(sideHeight, location int, orientation SideOrientation, sidePresence SidePresence) {
	store := *s
	sides, ok := store[sideHeight]
	if !ok {
		sides = make([]*Side, 0)
		store[sideHeight] = sides
	}

	start, end := location, location+1
	sideToRemove := -1
	var sideToEdit *Side
	for i, side := range sides {
		if side == nil {
			continue
		}
		if side.orientation != orientation || side.sidePresence != sidePresence {
			continue
		}
		if side.end == start {
			start = min(location, side.start)
			if sideToEdit != nil {
				sideToRemove = i - 1
			}
			sideToEdit = side
		}
		if side.start == end {
			end = max(location, side.end)

			if sideToEdit == nil {
				sideToEdit = side
			} else {
				sideToRemove = i
			}
		}
	}

	if sideToRemove != -1 {
		sides = slices.Delete(sides, sideToRemove, sideToRemove+1)
		store[sideHeight] = sides
	}

	if sideToEdit != nil {
		sideToEdit.start = start
		sideToEdit.end = end
	} else {
		newSide := Side{
			start:        start,
			end:          end,
			orientation:  orientation,
			sidePresence: sidePresence,
		}

		sides = append(sides, &newSide)
		store[sideHeight] = sides
	}
}
