package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Whojoo/AoC/2024/day4"
)

func main() {
	file, err := os.Open("input/day4.txt")
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

	firstResult := day4.HandleFirst(input)
	secondResult := day4.HandleSecond(input)

	fmt.Printf("First assignment result: %v\n", firstResult)
	fmt.Printf("Second assignment result: %v\n", secondResult)
}
