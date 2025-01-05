package main

import (
	"fmt"
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
	"time"

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
	responseChannels := make([]chan string, len(assignments))

	for i, assignment := range assignments {
		responseChannels[i] = make(chan string)

		go func() {
			input := shared.ReadInput("input/" + assignment.FileName())

			startTime := time.Now()
			part1 := assignment.Part1(input)
			part1Elapsed := time.Since(startTime)

			startTime = time.Now()
			part2 := assignment.Part2(input)
			part2Elapsed := time.Since(startTime)

			c := responseChannels[i]
			c <- getDay(i)
			c <- fmt.Sprintf("Part1: %d in %s", part1, part1Elapsed)
			c <- fmt.Sprintf("Part2: %d in %s", part2, part2Elapsed)
			c <- fmt.Sprintf("Total time: %s", part1Elapsed+part2Elapsed)

			close(c)
		}()
	}

	for _, responseChannel := range responseChannels {
		for response := range responseChannel {
			fmt.Println(response)
		}
		fmt.Println()
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
