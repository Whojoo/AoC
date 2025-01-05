package day11

import (
	"fmt"
	"github.com/Whojoo/AoC/2024/shared"
	"math"
	"strconv"
	"strings"
	"time"
)

type Assignment struct{}

func NewAssignment() *Assignment { return new(Assignment) }

func (Assignment) FileName() string { return "day11.txt" }

func (a Assignment) Handle(input []string, c chan<- string) {
	defer close(c)

	startTime := time.Now()

	part1, part2 := a.Part1(input), a.Part2(input)

	elapsed := time.Since(startTime)

	c <- "Day 11"
	c <- fmt.Sprintf("%d", part1)
	c <- fmt.Sprintf("%d", part2)
	c <- fmt.Sprintf("Took %s", elapsed)
}

func (Assignment) Part1(input []string) int {
	initialConfiguration := GenerateInitialConfiguration(input)
	finalConfiguration := PerformRulesOn(initialConfiguration, 25)
	return CountMembers(finalConfiguration)
}

func (Assignment) Part2(input []string) int {
	initialConfiguration := GenerateInitialConfiguration(input)
	finalConfiguration := PerformRulesOn(initialConfiguration, 75)
	return CountMembers(finalConfiguration)
}

func PerformRulesOn(initialConfiguration map[uint64]int, iterations int) map[uint64]int {
	configuration := initialConfiguration

	for range iterations {
		newConfiguration := make(map[uint64]int)

		for k, v := range configuration {
			for _, strat := range stoneBlinkStrategies {
				if !strat.IsApplicable(k) {
					continue
				}

				newValues := strat.Apply(k)

				for _, newValue := range newValues {
					newConfiguration[newValue] += v
				}

				break
			}
		}

		configuration = newConfiguration
	}

	return configuration
}

func CountMembers(configuration map[uint64]int) (sum int) {
	for _, v := range configuration {
		sum += v
	}

	return sum
}

func GenerateInitialConfiguration(input []string) map[uint64]int {
	fields := strings.Fields(input[0])
	values := shared.Project(fields, func(s string, _ int) uint64 {
		v, _ := strconv.ParseUint(s, 10, 64)
		return v
	})

	initialMap := make(map[uint64]int)

	for _, value := range values {
		initialMap[value]++
	}

	return initialMap
}

type StoneBlinkStrategy interface {
	IsApplicable(uint64) bool
	Apply(uint64) []uint64
}

var stoneBlinkStrategies = []StoneBlinkStrategy{
	NewZeroStoneBlinkStrategy(),
	NewEvenNumbersStoneBlinkStrategy(),
	NewTwentyTwentyFourStoneBlinkStrategy(),
}

type ZeroStoneBlinkStrategy struct{}

func NewZeroStoneBlinkStrategy() ZeroStoneBlinkStrategy   { return ZeroStoneBlinkStrategy{} }
func (ZeroStoneBlinkStrategy) IsApplicable(v uint64) bool { return v == 0 }
func (ZeroStoneBlinkStrategy) Apply(_ uint64) []uint64    { return []uint64{1} }

type EvenNumbersStoneBlinkStrategy struct{}

func NewEvenNumbersStoneBlinkStrategy() EvenNumbersStoneBlinkStrategy {
	return EvenNumbersStoneBlinkStrategy{}
}

func (s EvenNumbersStoneBlinkStrategy) IsApplicable(v uint64) bool {
	digits := s.getNumberOfDigits(v)
	return digits%2 == 0
}

func (s EvenNumbersStoneBlinkStrategy) Apply(v uint64) []uint64 {
	digits := s.getNumberOfDigits(v)
	splitAt := digits / 2
	power := uint64(math.Pow(10, float64(splitAt)))

	// Example: 2244
	// splitAt: 2, power: 100
	// left = 2244 / 100 = 22.44 => 22
	// right = 2244 - (22 * 100) = 2244 - 2200 = 44
	left := v / power
	right := v - (left * power)

	return []uint64{left, right}
}

func (s EvenNumbersStoneBlinkStrategy) getNumberOfDigits(v uint64) uint64 {
	if v == 0 {
		return uint64(1)
	}

	digits := uint64(1)
	comparison := uint64(10)

	for comparison <= v {
		comparison *= 10
		digits++
	}

	return digits
}

type TwentyTwentyFourStoneBlinkStrategy struct{}

func NewTwentyTwentyFourStoneBlinkStrategy() TwentyTwentyFourStoneBlinkStrategy {
	return TwentyTwentyFourStoneBlinkStrategy{}
}

func (TwentyTwentyFourStoneBlinkStrategy) IsApplicable(_ uint64) bool {
	return true
}

func (TwentyTwentyFourStoneBlinkStrategy) Apply(v uint64) []uint64 {
	return []uint64{v * 2024}
}
