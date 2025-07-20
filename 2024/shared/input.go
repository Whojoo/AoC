package shared

import (
	"bufio"
	"os"
)

func ReadInput(path string) (input []string) {
	file, err := os.Open(path)
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
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func ReadInputWithWeirdTokenPrevention(path string) (input []string) {
	input = ReadInput(path)

	if input[0][0] == 239 {
		// Weird file read thing
		input[0] = input[0][3:]
	}

	return input
}
