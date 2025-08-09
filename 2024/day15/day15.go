package day15

import (
	"fmt"
	"slices"
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
	crateLeft  = 14
	crateRight = 15
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
	world, directions, playerRow, playerColumn := parseInputPart2(input)

	for _, direction := range directions {
		var rowModifier, columnModifier int

		worldString := "Moving to "
		switch direction {
		case up:
			rowModifier = -1
			worldString += "^"
		case down:
			rowModifier = 1
			worldString += "v"
		case right:
			columnModifier = 1
			worldString += ">"
		case left:
			columnModifier = -1
			worldString += "<"
		}

		moved := recursiveMovePart2(world, playerRow, playerColumn, rowModifier, columnModifier)
		if moved {
			playerRow += rowModifier
			playerColumn += columnModifier
		}

		worldString += "\n"

		for row := range world {
			for _, tile := range world[row] {
				switch tile {
				case wall:
					worldString += "#"
				case emptySpace:
					worldString += "."
				case player:
					worldString += "@"
				case crateLeft:
					worldString += "["
				case crateRight:
					worldString += "]"
				}
			}

			worldString += "\n"
		}

		println(worldString)
		println()
	}

	score := 0

	for row := range world {
		for column, tile := range world[row] {
			if tile != crateLeft {
				continue
			}

			score += row*100 + column
		}
	}

	return score
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

type node struct{ x, y int }

func newNode(x, y int) *node {
	return &node{x, y}
}

func (n node) hash() int {
	return n.y*100 + n.x
}

type nodeStack []*node

func (stack *nodeStack) len() int {
	return len(*stack)
}

func (stack *nodeStack) peek() *node {
	dereferenced := *stack
	return dereferenced[stack.len()-1]
}

func (stack *nodeStack) push(newNode *node) {
	*stack = append(*stack, newNode)
}

func (stack *nodeStack) pop() *node {
	old := *stack
	popped := old[old.len()-1]
	*stack = old[:old.len()-1]
	return popped
}

type nodeSet map[int]*node

func (set *nodeSet) add(newNode *node) bool {
	if set.contains(newNode) {
		return false
	}

	(*set)[newNode.hash()] = newNode
	return true
}

func (set *nodeSet) contains(newNode *node) bool {
	_, found := (*set)[newNode.hash()]
	return found
}

func movePart2(world [][]int, startY, startX, yDir, xDir int) {
	stack, set := nodeStack{}, nodeSet{}

	stack.push(newNode(startX, startY))

	for stack.len() > 0 {
		popped := stack.pop()
		neighbourType := world[popped.y+yDir][popped.x+xDir]

		// Wall -> abort all checks, cannot move
		if neighbourType == wall {
			return
		}

		// Found a valid route, now continue other possible route checks
		if neighbourType == emptySpace {
			set.add(popped)
			continue
		}

		isCrate := neighbourType == crateLeft || neighbourType == crateRight

		if yDir != 0 && isCrate {
			if neighbourType == crateLeft {
				stack.push(newNode(popped.x, popped.y+yDir))
				stack.push(newNode(popped.x+1, popped.y+yDir))
			} else {
				stack.push(newNode(popped.x, popped.y+yDir))
				stack.push(newNode(popped.x-1, popped.y+yDir))
			}
		} else {
			stack.push(newNode(popped.x+xDir, popped.y+yDir))
		}
	}

	movables := []int{crateLeft, crateRight, player}
	for _, n := range set {
		for slices.Contains(movables, world[n.y-yDir][n.x-xDir]) {
			
		}
	}
}

func canCrateMove(world [][]int, leftX, leftY, rightX, rightY, yDir int) bool {
	leftNeighbour, rightNeighbour := world[leftY+yDir][leftX], world[rightY+yDir][rightX]
	leftCanMove := leftNeighbour == emptySpace || leftNeighbour == crateLeft || leftNeighbour == crateRight
	rightCanMove := rightNeighbour == emptySpace || rightNeighbour == crateLeft || rightNeighbour == crateRight
	return leftCanMove && rightCanMove
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

func parseInputPart2(input []string) (world [][]int, directions []int, playerRow, playerColumn int) {
	world = make([][]int, 0)
	directions = make([]int, 0)

	for row, line := range input {
		characters := strings.Split(line, "")

		if len(line) == 0 {
			continue
		}

		if characters[0] == "#" {
			world = append(world, make([]int, 0))
			for column, tile := range characters {
				switch tile {
				case "#":
					world[row] = append(world[row], wall)
					world[row] = append(world[row], wall)
				case "O":
					world[row] = append(world[row], crateLeft)
					world[row] = append(world[row], crateRight)
				case "@":
					world[row] = append(world[row], player)
					world[row] = append(world[row], emptySpace)
					playerRow = row
					playerColumn = column * 2
				case ".":
					world[row] = append(world[row], emptySpace)
					world[row] = append(world[row], emptySpace)
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
