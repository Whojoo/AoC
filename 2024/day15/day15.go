package day15

import (
	"fmt"
	"strings"
)

const (
	up         = 1
	left       = 2
	down       = 3
	right      = 4
	wall       = 10
	crate      = 11
	player     = 12
	emptySpace = 13
)

type Assignment struct{}

func NewAssignment() Assignment {
	return Assignment{}
}

func (a Assignment) FileName() string {
	return "day15.txt"
}

func (a Assignment) Part1(input []string) int {
	world, directions, playerRow, playerColumn := parseInput(input)

	for _, direction := range directions {
		var rowModifier, columnModifier int

		//worldString := "Moving to "
		switch direction {
		case up:
			rowModifier = -1
			//worldString += "^"
		case down:
			rowModifier = 1
			//worldString += "v"
		case right:
			columnModifier = 1
			//worldString += ">"
		case left:
			columnModifier = -1
			//worldString += "<"
		}

		moved := recursiveMove(world, playerRow, playerColumn, rowModifier, columnModifier)
		if moved {
			playerRow += rowModifier
			playerColumn += columnModifier
		}

		//worldString += "\n"

		//for row := range world {
		//	for _, tile := range world[row] {
		//		switch tile {
		//		case wall:
		//			worldString += "#"
		//		case emptySpace:
		//			worldString += "."
		//		case player:
		//			worldString += "@"
		//		case crate:
		//			worldString += "O"
		//		}
		//	}
		//
		//	worldString += "\n"
		//}

		//println(worldString)
		//println()
	}

	score := 0

	for row := range world {
		for column, tile := range world[row] {
			if tile != crate {
				continue
			}

			score += row*100 + column
		}
	}

	return score
}

func (a Assignment) Part2(input []string) int {
	return len(input)
}

func recursiveMove(world [][]int, startY, startX, yDir, xDir int) (moved bool) {
	neighbourType := world[startY+yDir][startX+xDir]

	if neighbourType == wall {
		return
	}

	if neighbourType == emptySpace {
		world[startY][startX], world[startY+yDir][startX+xDir] = world[startY+yDir][startX+xDir], world[startY][startX]
		return true
	}

	recursiveMove(world, startY+yDir, startX+xDir, yDir, xDir)
	neighbourType = world[startY+yDir][startX+xDir]

	if neighbourType == emptySpace {
		world[startY][startX], world[startY+yDir][startX+xDir] = world[startY+yDir][startX+xDir], world[startY][startX]
		moved = true
	}

	return moved
}

func parseInput(input []string) (world [][]int, directions []int, playerRow, playerColumn int) {
	world = make([][]int, 0)
	directions = make([]int, 0)

	for row, line := range input {
		characters := strings.Split(line, "")

		if len(line) == 0 {
			continue
		}

		if characters[0] == "#" {
			world = append(world, make([]int, len(characters)))
			for column, tile := range characters {
				switch tile {
				case "#":
					world[row][column] = wall
				case "O":
					world[row][column] = crate
				case "@":
					world[row][column] = player
					playerRow = row
					playerColumn = column
				case ".":
					world[row][column] = emptySpace
				default:
					fmt.Printf("Found weird world character %s\n", tile)
				}
			}
		} else {
			for _, directionCharacter := range characters {
				switch directionCharacter {
				case "^":
					directions = append(directions, up)
				case "<":
					directions = append(directions, left)
				case "v":
					directions = append(directions, down)
				case ">":
					directions = append(directions, right)
				default:
					fmt.Printf("Found weird direction character %s\n", directionCharacter)
				}
			}
		}
	}

	return world, directions, playerRow, playerColumn
}
