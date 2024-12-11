package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Whojoo/AoC/2024/day1"
)

func main() {
	fileName := day1.GetFileName()
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	assignment := day1.CreateAssignment(input)
	firstResult := assignment.HandleFirst()
	secondResult := assignment.HandleSecond()

	fmt.Printf("First assignment result: %v\n", firstResult)
	fmt.Printf("Second assignment result: %v\n", secondResult)
}
