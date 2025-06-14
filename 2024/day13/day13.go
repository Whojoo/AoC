package day13

import (
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
	part2Max         = math.MaxInt
	buttonACost      = 3
	buttonBCost      = 1
)

func (a Assignment) Part1(input []string) int {
	games := getGames(input)
	totalCost := 0

	for _, game := range games {
		solution, success := game.FindSolution(maxButtonPresses)
		if !success {
			continue
		}

		totalCost += solution.ButtonAPresses*buttonACost + solution.ButtonBPresses*buttonBCost
	}

	return totalCost
}

func (a Assignment) Part2(input []string) int {
	games := getGames(input)
	totalCost := 0
	const prizeAddition = 10_000_000_000_000

	for _, game := range games {
		game.PrizeVector = Vector{x: game.PrizeVector.x + prizeAddition, y: game.PrizeVector.y + prizeAddition}
		solution, success := game.FindSolution(part2Max)
		if !success {
			continue
		}

		totalCost += solution.ButtonAPresses*buttonACost + solution.ButtonBPresses*buttonBCost
	}

	return totalCost
}

type Game struct {
	ButtonAVector, ButtonBVector, PrizeVector Vector
}

type GameSolution struct {
	ButtonAPresses, ButtonBPresses int
}

func NewGameSolution(buttonAPresses, buttonBPresses int) GameSolution {
	return GameSolution{
		ButtonAPresses: buttonAPresses,
		ButtonBPresses: buttonBPresses,
	}
}

func (game Game) FindSolution(maxPresses int) (solution GameSolution, found bool) {
	// Cheated, but I do not understand this math and just wanted to move on
	determinant := (game.ButtonAVector.x * game.ButtonBVector.y) - (game.ButtonAVector.y * game.ButtonBVector.x)

	//  No way to reach the prize
	if determinant == 0 {
		return GameSolution{}, false
	}

	buttonAPresses := (game.ButtonBVector.y*game.PrizeVector.x - game.ButtonBVector.x*game.PrizeVector.y) / determinant
	buttonBPresses := (game.ButtonAVector.x*game.PrizeVector.y - game.ButtonAVector.y*game.PrizeVector.x) / determinant

	if buttonAPresses < 0 || buttonBPresses < 0 || math.Mod(buttonAPresses, 1) != 0 || math.Mod(buttonBPresses, 1) != 0 {
		return GameSolution{}, false
	}

	exactButtonAPresses := int(math.Ceil(buttonAPresses))
	exactButtonBPresses := int(math.Ceil(buttonBPresses))

	if exactButtonAPresses > maxPresses || exactButtonBPresses > maxPresses {
		return GameSolution{}, false
	}

	return NewGameSolution(exactButtonAPresses, exactButtonBPresses), true
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

func (vector Vector) Reversed() Vector {
	return vector.Scale(-1)
}

func (vector Vector) Direction() float64 {
	return math.Tan(vector.y / vector.x)
}

// func (vector Vector) Normalize() Vector {
//	scalar := 1 / vector.Length()
//	return vector.Scale(scalar)
// }
