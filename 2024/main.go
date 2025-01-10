package main

import (
	"fmt"
	"time"

	"github.com/Whojoo/AoC/2024/day1"
	"github.com/Whojoo/AoC/2024/day10"
	"github.com/Whojoo/AoC/2024/day11"
	"github.com/Whojoo/AoC/2024/day12"
	"github.com/Whojoo/AoC/2024/day2"
	"github.com/Whojoo/AoC/2024/day3"
	"github.com/Whojoo/AoC/2024/day4"
	"github.com/Whojoo/AoC/2024/day8"
	"github.com/Whojoo/AoC/2024/day9"
	"github.com/Whojoo/AoC/2024/shared"

	"github.com/Whojoo/AoC/2024/day6"
)

func main() {
	assignments := []shared.Assignment{
		day1.GetAssignment(),
		day2.GetAssignment(),
		day3.GetAssignment(),
		day4.GetAssignment(),
		day6.GetAssignment(),
		day8.GetAssignment(),
		day9.GetAssignment(),
		day10.NewAssignment(),
		day11.NewAssignment(),
		day12.NewAssignment(),
	}

	for i, assignment := range assignments {
		input := shared.ReadInput("input/" + assignment.FileName())

		startTime := time.Now()
		part1 := assignment.Part1(input)
		part1Elapsed := time.Since(startTime)

		startTime = time.Now()
		part2 := assignment.Part2(input)
		part2Elapsed := time.Since(startTime)

		fmt.Println(getDay(i))
		fmt.Printf("Part1: %d in %s\n", part1, part1Elapsed)
		fmt.Printf("Part2: %d in %s\n", part2, part2Elapsed)
		fmt.Printf("Total time: %s\n\n", part1Elapsed+part2Elapsed)
	}
}

var skippedDays = []int{5, 7}

func getDay(index int) string {
	day := index + 1
	for _, skippedDay := range skippedDays {
		if day >= skippedDay {
			day++
		}
	}

	return fmt.Sprintf("Day %d", day)
}
