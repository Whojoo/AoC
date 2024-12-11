package day1

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day1 struct {
	input []string
}

func GetFileName() string {
	return "day1.txt"
}

func CreateAssignment(input []string) Day1 {
	assignment := Day1{
		input: input,
	}

	return assignment
}

func (assignment Day1) HandleFirst() int {
	// Change to 2 int slices
	left, right := createIntSlices(assignment.input)

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

func (assignment Day1) HandleSecond() int {
	// Change to 2 int slices
	left, right := createIntSlices(assignment.input)

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