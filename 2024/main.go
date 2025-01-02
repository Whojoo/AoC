package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/Whojoo/AoC/2024/day1"
	"github.com/Whojoo/AoC/2024/day2"
	"github.com/Whojoo/AoC/2024/day3"
	"github.com/Whojoo/AoC/2024/day4"
	"github.com/Whojoo/AoC/2024/shared"

	"github.com/Whojoo/AoC/2024/day6"
)

func main() {
	assignments := []shared.Assignment{
		day1.GetAssignment(),
		day2.GetAssignment(),
		day3.GetAssignment(),
		day4.GetAssignment(),
		day6.GetAssignment2(),
	}
	responseChannels := make([]chan string, len(assignments))

	for i, assignment := range slices.Backward(assignments) {
		responseChannels[i] = make(chan string)

		go func() {
			input := getInput(assignment.FileName())
			assignment.Handle(input, responseChannels[i])
		}()
	}

	for _, responseChannel := range slices.Backward(responseChannels) {
		for response := range responseChannel {
			fmt.Println(response)
		}
		fmt.Println()
	}
}

func getInput(fileName string) []string {
	file, err := os.Open("input/" + fileName)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
