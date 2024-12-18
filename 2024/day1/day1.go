package day1

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Assignment struct{}

func GetAssignment() Assignment {
	return Assignment{}
}

func (Assignment) Handle(input []string, c chan<- string) {
	startTime := time.Now()
	first := HandleFirst(input)
	firstTime := time.Since(startTime)

	startTime = time.Now()
	second := HandleSecond(input)
	secondTime := time.Since(startTime)

	c <- "Day 1"
	c <- fmt.Sprintf("First result: %d in %s", first, firstTime)
	c <- fmt.Sprintf("Second result: %d in %s", second, secondTime)

	close(c)
}

func (Assignment) FileName() string {
	return "day1.txt"
}

func HandleFirst(input []string) int {
	// Change to 2 int slices
	left, right := createIntSlices(input)

	// Order slices
	slices.Sort(left)
	slices.Sort(right)

	// Loop and sum
	sum := float64(0)

	for i := range left {
		sum += math.Abs(float64(left[i]) - float64(right[i]))
	}

	return int(sum)
}

func HandleSecond(input []string) int {
	// Change to 2 int slices
	left, right := createIntSlices(input)

	// Map right to distinct entries with number of occurrences
	rightMap := make(map[int]int)
	for _, value := range right {
		rightMap[value]++
	}

	// Use left as key, multiply it by the occurrences
	sum := 0
	for _, value := range left {
		sum += value * rightMap[value]
	}

	return sum
}

func createIntSlices(input []string) ([]int, []int) {
	left, right := make([]int, len(input)), make([]int, len(input))

	for i, line := range input {
		fields := strings.Fields(line)

		left[i], _ = strconv.Atoi(fields[0])
		right[i], _ = strconv.Atoi(fields[1])
	}

	return left, right
}
