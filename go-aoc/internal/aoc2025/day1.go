package aoc2025

import (
	"strconv"
)

const startingPosition int = 50
const numOfOptions int = 100

func Day1(in []string) (int, int, error) {
	p1, err := part1(in)
	if err != nil {
		return 0, 0, err
	}

	return p1, 0, nil
}

func part1(input []string) (int, error) {
	o, err := mapToOrders(input)

	if err != nil {
		return 0, err
	}

	r := 0
	p := startingPosition

	for _, v := range o {
		p = p + v

		if p%numOfOptions == 0 {
			r = r + 1
		}
	}

	return r, nil
}

func mapToOrders(in []string) ([]int, error) {
	out := make([]int, len(in))

	for i, v := range in {
		n, err := strconv.Atoi(v[1:])
		if err != nil {
			return []int{}, err
		}

		if v[0] == 'L' {
			n = n * -1
		}

		out[i] = n
	}

	return out, nil
}
