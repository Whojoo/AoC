package day13

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Assignment struct{}

func NewAssignment() Assignment {
	return Assignment{}
}

func (a Assignment) FileName() string {
	return "day13.txt"
}

const (
	maxButtonPresses = 100
)

func (a Assignment) Part1(input []string) int {
	games := getGames(input)

	for _, game := range games {
		fmt.Println(game)
	}

	return len(games)
}

func (a Assignment) Part2(input []string) int {
	return len(input)
}

type Game struct {
	ButtonAVector, ButtonBVector, PrizeVector Vector
}

type GameSolution struct {
	DirectionVector                Vector
	ButtonAPresses, ButtonBPresses int
}

func NewGameSolution(directionVector Vector, buttonAPresses, buttonBPresses int) GameSolution {
	return GameSolution{
		DirectionVector: directionVector,
		ButtonAPresses:  buttonAPresses,
		ButtonBPresses:  buttonBPresses,
	}
}

func (game Game) FindSolution() (solution GameSolution, found bool) {
	if !game.hasSolution() {
		return solution, false
	}

	solution, found = game.findEarlySolution()

	if found {
		return solution, found
	}

	var buttonAPressed, buttonBPressed int
	currentVector := game.ButtonBVector
	prizeDirection := game.PrizeVector.Direction()
	aDirection, bDirection := game.ButtonAVector, game.ButtonBVector

	left, right := &game.ButtonAVector, &game.ButtonBVector
	fmt.Println(left, right)
	if aDirection.Direction() < bDirection.Direction() {
		left, right = right, left
		fmt.Println(left, right)
	}

	for currentDirection := currentVector.Direction(); currentDirection != prizeDirection &&
		(buttonAPressed+buttonBPressed) <= maxButtonPresses; {
	}

	return GameSolution{}, true
}

func (game Game) findEarlySolution() (solution GameSolution, found bool) {
	prizeDirection := game.PrizeVector.Direction()
	aDirection, bDirection := game.ButtonAVector.Direction(), game.ButtonBVector.Direction()

	if bDirection == prizeDirection {
		solution = NewGameSolution(game.ButtonBVector, 0, 1)
		return solution, true
	}

	if aDirection == prizeDirection {
		solution = NewGameSolution(game.ButtonAVector, 1, 0)
		return solution, true
	}

	return solution, false
}

func (game Game) hasSolution() bool {
	prizeDirection := game.PrizeVector.Direction()
	aDirection, bDirection := game.ButtonAVector.Direction(), game.ButtonBVector.Direction()

	validWithALeft := aDirection >= prizeDirection && prizeDirection >= bDirection
	validWithBLeft := bDirection >= prizeDirection && prizeDirection >= aDirection

	return validWithALeft || validWithBLeft
}

func getGames(input []string) []Game {
	// Each case has 2 button lines and a prize line. All cases but last also have a blank line
	const linesPerGame = 4
	numberOfGames := (len(input) + 1) / linesPerGame
	games := make([]Game, numberOfGames)[:0]

	vectorRegex := regexp.MustCompile(`[A-Za-z\s=:+]*(\d+),[A-Za-z\s=:+]*(\d+)`)

	for i := 0; i < len(input); i += linesPerGame {
		// Button A = n, Button B = n+1, Prize = n+2
		buttonAMatches := vectorRegex.FindStringSubmatch(input[i])
		buttonBMatches := vectorRegex.FindStringSubmatch(input[i+1])
		prizeMatches := vectorRegex.FindStringSubmatch(input[i+2])

		game := Game{
			ButtonAVector: NewVectorFromMatches(buttonAMatches),
			ButtonBVector: NewVectorFromMatches(buttonBMatches),
			PrizeVector:   NewVectorFromMatches(prizeMatches),
		}
		games = append(games, game)
	}

	return games
}

type Vector struct{ x, y float64 }

func NewVectorFromMatches(matches []string) Vector {
	x, _ := strconv.ParseFloat(matches[1], 64)
	y, _ := strconv.ParseFloat(matches[2], 64)
	return Vector{x, y}
}

func (vector Vector) X() float64 { return vector.x }
func (vector Vector) Y() float64 { return vector.y }

func (vector Vector) Add(otherVector Vector) Vector {
	return Vector{vector.x + otherVector.x, vector.y + otherVector.y}
}

func (vector Vector) Length() float64 {
	return math.Sqrt(vector.x*vector.x + vector.y*vector.y)
}

func (vector Vector) Scale(scalar float64) Vector {
	return Vector{vector.x * scalar, vector.y * scalar}
}

func (vector Vector) Direction() float64 {
	return math.Tan(vector.y / vector.x)
}

// func (vector Vector) Normalize() Vector {
//	scalar := 1 / vector.Length()
//	return vector.Scale(scalar)
// }
