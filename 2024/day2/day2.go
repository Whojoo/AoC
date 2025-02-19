package day2

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Assignment struct{}

func (a Assignment) Part1(input []string) int {
	// Create int slices
	reports := createIntSlices(input)

	sum := 0
	for _, report := range reports {
		sum += checkReport(report)
	}

	return sum
}

func (a Assignment) Part2(input []string) int {
	// Create int slices
	reports := createIntSlices(input)

	sum := 0
	for _, report := range reports {
		sum += checkReportWithSafetyMargin(report)
	}

	return sum
}

func GetAssignment() Assignment {
	return Assignment{}
}

func (a Assignment) Handle(input []string, c chan<- string) {
	startTime := time.Now()
	first := a.Part1(input)
	firstTime := time.Since(startTime)

	startTime = time.Now()
	second := a.Part2(input)
	secondTime := time.Since(startTime)

	c <- "Day 2"
	c <- fmt.Sprintf("First result: %d in %s", first, firstTime)
	c <- fmt.Sprintf("Second result: %d in %s", second, secondTime)

	close(c)
}

func (Assignment) FileName() string {
	return "day2.txt"
}

type direction int

var (
	up   direction = 1
	down direction = -1
)

func checkReport(report []int) int {
	if len(report) < 2 || report[0] == report[1] {
		return 0
	}

	reportDirection := up
	if report[0] > report[1] {
		reportDirection = down
	}

	for i := range report {
		if i+1 == len(report) {
			break
		}

		diff := report[i+1] - report[i]

		if diff == 0 {
			return 0
		}

		isRightDirection := direction(diff/abs(diff)) == reportDirection
		isSteadyDiff := abs(diff) >= 1 && abs(diff) <= 3

		if !isRightDirection || !isSteadyDiff {
			return 0
		}
	}

	return 1
}

func abs(input int) int {
	return int(math.Abs(float64(input)))
}

func checkReportWithSafetyMargin(report []int) int {
	initialResult := checkReport(report)

	if initialResult > 0 {
		return initialResult
	}

	safetyResults := make([]int, len(report))

	var wg sync.WaitGroup

	for i := range report {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			slicedReport := slices.Clone(report)
			slicedReport = slices.Delete(slicedReport, index, index+1)
			safetyResults[index] = checkReport(slicedReport)
		}(i)
	}

	wg.Wait()

	for _, result := range safetyResults {
		if result == 1 {
			return 1
		}
	}

	return 0
}

func createIntSlices(input []string) [][]int {
	left := make([][]int, len(input))

	for i, line := range input {
		fields := strings.Fields(line)
		left[i] = make([]int, len(fields))

		for j, field := range fields {
			left[i][j], _ = strconv.Atoi(field)
		}
	}

	return left
}
