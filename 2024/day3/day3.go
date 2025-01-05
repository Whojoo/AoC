package day3

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type Assignment struct{}

func (a Assignment) Part1(input []string) int {
	// Create single string
	mergedInput := reduce(input)

	// Regex rules the world
	matches := firstRegex.FindAllStringSubmatch(mergedInput, -1)

	sum := 0
	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		sum += left * right
	}

	return sum
}

func (a Assignment) Part2(input []string) int {
	mergedInput := reduce(input)
	enabled := true
	sum := 0
	matched := secondRegex.FindAllStringSubmatch(mergedInput, -1)

	for _, match := range matched {
		if match[0] == "don't()" {
			enabled = false
			continue
		}
		if match[0] == "do()" {
			enabled = true
			continue
		}

		if enabled {
			left, _ := strconv.Atoi(match[2])
			right, _ := strconv.Atoi(match[3])
			sum += left * right
		}
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

	c <- "Day 3"
	c <- fmt.Sprintf("First result: %d in %s", first, firstTime)
	c <- fmt.Sprintf("Second result: %d in %s", second, secondTime)

	close(c)
}

func (Assignment) FileName() string {
	return "day3.txt"
}

var (
	firstRegex  = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	secondRegex = regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)
)

func reduce(input []string) string {
	var output string

	for _, line := range input {
		output += line
	}

	return output
}
